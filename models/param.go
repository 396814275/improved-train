package models

const (
	OrderByTime  = "time"
	OrderByScore = "score"
)

// ParamSignUp 用户注册结构体
type ParamSignUp struct {
	//三个字段都必须要有值 required
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

// ParamLogin 用户登录结构体
type ParamLogin struct {
	//两个字段都必须要有值 required
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// ParamVoteData 获取投票参数
type ParamVoteData struct {
	//UserID
	Uid       int64 `db:"post_id"`                                             //标记用户->帖子的ID
	PostID    int64 `json:"post_id,string" binding:"required"`                 //帖子ID
	Direction int   `json:"direction,string"  binding:"required,oneof=0 1 -1"` //赞成or反对
}

// ParamPostList 获取帖子列表参数
type ParamPostList struct {
	Page  int    `form:"page" json:"page"`
	Size  int    `form:"size" json:"size"`
	Order string `form:"order" json:"order"`
}
type ParamCommunityPostList struct {
	*ParamPostList
	CommunityID int64 `json:"community_id" form:"community_id"`
}
