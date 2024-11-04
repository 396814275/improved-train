package logic

import (
	"strconv"
	"web2/dao/redis"
	"web2/models"
)

//投票功能
/*1.用户投票
2.投票的几种情况：
direction=1：
  1.未投票，现在投票
  2.反对，先赞成
direction=0:
  1.赞成，现取消
  2.反对，现取消
direction=-1:
  1.未投票，现投票
  2.赞成，现反对
*/

func PostVote(p *models.ParamVoteData, userID int64) error {
	return redis.VoteForPost(strconv.Itoa(int(userID)), strconv.FormatInt(p.PostID, 20), float64(p.Direction))

}
