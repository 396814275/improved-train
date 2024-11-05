package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web2/logic"
	"web2/models"
)

func CreatePostHandler(c *gin.Context) {
	//1.获取参数及参数的校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("create post with invalid param:%v", zap.Error(err))
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
func GetPostDetailHandler(c *gin.Context) {
	//	获取参数
	pidStr := c.Param("id")
	pid, err := strconv.ParseInt(pidStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetail failed with a invalid param", zap.Error(err))
		ResponseError(c, CodeInvalidParam)
	}
	//	根据id取出帖子数据（查数据库）
	data, err := logic.GetPostByID(pid)
	if err != nil {
		zap.L().Error("logic.GetPostDetail failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//返回响应
	ResponseSuccess(c, data)
}
func GetPostListHandler(c *gin.Context) {
	//获取分页参数
	page, size := GetPageInfo(c)

	//	获取数据
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	//	返回响应
	ResponseSuccess(c, data)
}
