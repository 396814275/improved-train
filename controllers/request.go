package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const ContextUserIDKey = "UserID"

var ErrorUserNotLogin = errors.New("用户未登录")

func GetCurrentUser(c *gin.Context) (UserID int64, err error) {
	uid, ok := c.Get("UserID")
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	UserID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	return

}
