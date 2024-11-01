package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"net/http"
	"web2/dao/mysql"
	"web2/logic"
	"web2/models"
)

// SignUpHandler 处理注册请求的函数
func SignUpHandler(c *gin.Context) {
	//	1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Signup with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return

	}
	//手动对请求参数进行详细的业务规矩校验
	//if len(p.Username) == 0 || len(p.Password) == 0 || len(p.RePassword) == 0 || p.RePassword != p.Password {
	//	zap.L().Error("Signup with invalid param")
	//	c.JSON(http.StatusOK, gin.H{
	//		"msg": "请求参数有误",
	//	})
	//	return
	//}

	//	2.业务处理
	if err := logic.SignUp(p); err != nil {
		zap.L().Error("signup failed err:", zap.Error(err))
		if errors.Is(err, mysql.ErrUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServeBusy)
		return
	}
	//	3.返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 处理登录请求的函数
func LoginHandler(c *gin.Context) {
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("Login with invalid param", zap.Error(err))
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		c.JSON(http.StatusOK, gin.H{
			"msg": removeTopStruct(errs.Translate(trans)),
		})
		return

	}
	//2.业务逻辑处理
	user, err := logic.Login(p)
	if err != nil {
		zap.L().Error("Login failed err:", zap.String("username:", p.Username), zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//	3.返回响应
	ResponseSuccess(c, gin.H{
		"user_id":  fmt.Sprintf("%d", user.UserId),
		"username": user.Username,
		"token":    user.Token,
	})
}
