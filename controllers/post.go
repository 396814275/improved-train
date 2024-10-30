package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"web2/logic"
	"web2/models"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invalid param")
		ResponseError(c, CodeInvalidParam)
		return
	}
	//从c 取到当前用户的ID
	UserID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
		zap.L().Error("GetCurrentUser failed", zap.Error(err))
	}
	p.AuthorID = UserID
	//	2.创建帖子
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("Logic.CreatePost failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}
