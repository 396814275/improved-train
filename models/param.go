package models

// 用户注册结构体
type ParamSignUp struct {
	//三个字段都必须要有值 required
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"repassword" binding:"required,eqfield=Password"`
}

// 用户登录结构体
type ParamLogin struct {
	//两个字段都必须要有值 required
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
