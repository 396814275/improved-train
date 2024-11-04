package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"web2/logic"
	"web2/models"
)

//type VoteData struct {
//	//UserID
//	PostID     int64 `json:"post_id,string"`   //帖子ID
//	Directtion int8  `json:"direction,string"` //赞成or反对
//}

func PostVoteHandler(c *gin.Context) {
	//	参数校验
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		errData := removeTopStruct(errs.Translate(trans))
		ResponseErrorWithMsg(c, CodeInvalidParam, errData)
		return
	}
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CodeNeedLogin)
	}
	if err := logic.PostVote(p, userID); err != nil {
		zap.L().Error("logic.PostVote failed", zap.Error(err))
		ResponseError(c, CodeServeBusy)
		return
	}
	ResponseSuccess(c, nil)
}
