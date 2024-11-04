package redis

// redis key 注意使用命名空间的方式，方便查询的拆分
const (
	KeyPrefix              = "blueball:"
	KeyPostTimeZset        = "post:time"   //zset:帖子及发表时间
	KeyPostScoreZset       = "post:score"  //zset:帖子及投票的分数
	KeyPostScoreVoteZsetPF = "post:voted:" //zset:记录用户及投票类型；参数是Postid
)
