package redis

import (
	"context"
	"errors"
	"github.com/redis/go-redis/v9"
	"math"
	"strconv"
	_ "strconv"
	"time"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePerVote     = 432 // 每一票值多少分
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过")
	ErrVoteRepeated   = errors.New("不允许重复投票")
)

func CreatePost(postId, communityID int64) error {

	pipeline := client.TxPipeline()
	pipeline.ZAdd(context.TODO(), getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	pipeline.ZAdd(context.TODO(), getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})

	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
	pipeline.SAdd(context.TODO(), cKey, postId)

	_, err := pipeline.Exec(context.TODO())

	return err
}

func VoteForPost(userID, postID string, value float64) error {
	// 1. 判断投票限制
	// 去redis取帖子发布时间
	postTime := client.ZScore(context.TODO(), getRedisKey(KeyPostTimeZSet), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpire
	}
	// 2和3需要放到一个pipeline事务中操作

	// 2. 更新贴子的分数
	// 先查当前用户给当前帖子的投票记录
	ov := client.ZScore(context.TODO(), getRedisKey(KeyPostVotedZSetPF+postID), userID).Val()

	// 更新：如果这一次投票的值和之前保存的值一致，就提示不允许重复投票
	if value == ov {
		return ErrVoteRepeated
	}
	var op float64
	if value > ov {
		op = 1
	} else {
		op = -1
	}
	diff := math.Abs(ov - value) // 计算两次投票的差值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(context.TODO(), getRedisKey(KeyPostScoreZSet), op*diff*scorePerVote, postID)

	// 3. 记录用户为该贴子投票的数据
	if value == 0 {
		pipeline.ZRem(context.TODO(), getRedisKey(KeyPostVotedZSetPF+postID), userID)
	} else {
		pipeline.ZAdd(context.TODO(), getRedisKey(KeyPostVotedZSetPF+postID), redis.Z{
			Score:  value, // 赞成票还是反对票
			Member: userID,
		})
	}
	_, err := pipeline.Exec(context.TODO())
	return err
}
