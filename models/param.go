package models

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
type ParamVoteData struct {
	//UserID
	PostID    int64 `json:"post_id,string" binding:"required"`                 //帖子ID
	Direction int8  `json:"direction,string"  binding:"required,oneof=1 0 -1"` //赞成or反对
}
