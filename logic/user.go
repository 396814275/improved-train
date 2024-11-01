package logic

import (
	"web2/dao/mysql"
	"web2/models"
	"web2/pkg/jwt"
	"web2/pkg/snowflake"
)

//存放业务逻辑的代码，可能会多次调用dao层

func SignUp(p *models.ParamSignUp) (err error) {
	//1.判断用户名是否已被注册

	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}

	//2.生成UID
	userID := snowflake.GenID()
	//构造User实例
	u := &models.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password}

	//3.密码加密

	//4.数据存进数据库
	err = mysql.InsertUser(u)
	if err != nil {
		return err
	}

	return
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	//传递的是指针
	if err := mysql.Login(user); err != nil {
		return nil, nil
	}
	token, err := jwt.GenToken(user.UserId, user.Username)
	if err != nil {
		return nil, err
	}
	user.Token = token
	return
}
