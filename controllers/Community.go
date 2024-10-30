package controllers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
	"web2/logic"
)

//--跟社区相关的--

func CommunityHandler(c *gin.Context) {
	//拉取社区数据（community-id,community-name）以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情
func CommunityDetailHandler(c *gin.Context) {
	communityID := c.Param("id")
	id, err := strconv.ParseInt(communityID, 10, 64)
	if err != nil {
		ResponseError(c, CodeInvalidParam)
		return
	}
	//拉取社区数据（community-id,community-name）以列表的形式返回
	data, err := logic.GetCommunityDetailList(id)
	if err != nil {
		zap.L().Error("logic.GetCommunityList failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, data)
}
