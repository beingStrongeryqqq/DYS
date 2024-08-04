package models

// 定义请求的参数结构体

const (
	OrderTime  = "time"
	OrderScore = "score"
)

type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
}

type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamVoteData struct {
	PostID    string `json:"post_id" binding:"required"`
	Direction int8   `json:"direction,string" binding:"oneof=1 0 -1"`
}

type ParamPostList struct {
	Page        int64  `form:"page"`
	Size        int64  `form:"size"`
	Order       string `form:"order"`
	CommunityID int64  `form:"community_id"`
}

/*
type ParamCommunityPostList struct {
	ParamPostList
	CommunityID int64 `form:"community_id"`
}
*/
