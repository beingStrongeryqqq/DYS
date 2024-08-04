package redis

const (
	Prefix             = "dys:"   //zset
	KeyPostTimeZSet    = "post:time"   //zset
	KeyPostScoreZSet   = "post:score"  //zset
	KeyPostVotedZSetPF = "post:voted:" //zset
	KeyCommunitySetPF  = "community:"  //set
)

func getRedisKey(key string) string {
	return Prefix + key
}
