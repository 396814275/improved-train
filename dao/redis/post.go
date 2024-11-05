package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"web2/models"
)

func GetPostInOrder(p *models.ParamPostList) ([]string, error) {
	//	从redis中获取id
	//根据用户请求中携带的参数确定要查询的redis key
	key := getRedisKey(KeyPostTimeZset)
	if p.Order == models.OrderByScore {
		key = getRedisKey(KeyPostScoreZset)
	}
	//确定查询的索引起始点
	start := (p.Page - 1) * p.Size
	end := start + p.Size - 1
	//ZRevRange
	ids, err := rdb.ZRevRange(context.Background(), key, start, end).Result()
	if err != nil {
		return nil, err
	}
	return ids, nil
}

// GetPostVoteData 根据ids查询每篇帖子的票数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostScoreVoteZsetPF + id)
	//	//查找key中分数是1的元素的数量->统计每篇帖子的赞成票的数量
	//	v1 := rdb.ZCount(context.Background(), key, "1", "1").Val()
	//	data = append(data, v1)
	//}
	//使用pipeline一次发送多条命令，减少rtt
	pipeline := rdb.Pipeline()
	data = make([]int64, 0, len(ids))

	for _, id := range ids {
		key := getRedisKey(KeyPostScoreVoteZsetPF + id)
		pipeline.ZCount(context.Background(), key, "1", "1")
	}
	cmders, err := pipeline.Exec(context.Background())
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}

	return
}

// GetCommunityPostIDsInOrder 按社区查询ids
//func GetCommunityPostIDsInOrder(orderKey string, communityID, page, size int64) ([]string, error) {
//	//使用zinterstore 把分区的帖子set和帖子分数的zset 生成一个新的zset
//	//针对新的zset 按之前的逻辑取数据
//
//	//社区的key
//	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityID)))
//	//利用缓存key减少zinterstore执行的次数
//	key := orderKey + strconv.Itoa(int(communityID))
//	if rdb.Exists(context.Background(), orderKey).Val() < 1 {
//
//		rdb.ZInterStore(context.Background(), key, &redis.ZStore{
//			Aggregate: "MAX",
//		})
//	}
//	//	从redis中获取id
//	//根据用户请求中携带的参数确定要查询的redis key
//	key := getRedisKey(KeyPostTimeZset)
//	if p.Order == models.OrderByScore {
//		key = getRedisKey(KeyPostScoreZset)
//	}
//	//确定查询的索引起始点
//	start := (p.Page - 1) * p.Size
//	end := start + p.Size - 1
//	//ZRevRange
//	ids, err := rdb.ZRevRange(context.Background(), key, start, end).Result()
//	if err != nil {
//		return nil, err
//	}
//	return ids, nil
//}
