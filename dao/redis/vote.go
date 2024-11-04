package redis

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"math"
	"time"
)

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

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432
)

var (
	ErrVoteTimeExpired = errors.New("vote time expired")
)

func VoteForPost(userID string, postID string, value float64) (err error) {
	//  1.接受参数并校验限制

	postTime := rdb.ZScore(context.Background(), getRedisKey(KeyPostTimeZset), postID).Val()
	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		return ErrVoteTimeExpired
	}
	//	2.更新
	//先查当前用户给当前帖子的投票记录
	ov := rdb.ZScore(context.Background(), getRedisKey(KeyPostScoreVoteZsetPF+postID), userID).Val()
	diff := math.Abs(ov - value) //计算两次操作的差值的绝对值
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir -= 1
	}
	_, err = rdb.ZIncrBy(context.Background(), getRedisKey(KeyPostScoreZset), dir*diff*scorePreVote, postID).Result()
	if err != nil {
		return err
	}
	//	3.记录用户的操作
	if value == 0 {
		_, err = rdb.ZRem(context.Background(), getRedisKey(KeyPostScoreZset+postID), postID).Result()
	} else {
		_, err = rdb.ZAdd(context.Background(), getRedisKey(KeyPostScoreVoteZsetPF+postID), &redis.Z{
			Score:  value,
			Member: userID,
		}).Result()
	}
	return err
}
