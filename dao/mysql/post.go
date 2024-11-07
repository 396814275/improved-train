package mysql

import (
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"strings"

	"web2/models"
)

func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(
    post_id,author_id,community_id,title,content)
    values (?,?,?,?,?)`
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)
	if err != nil {
		zap.L().Error("post insert to databases failed")
	}
	return
}
func GetPostByID(pid int64) (post *models.Post, err error) {
	post = new(models.Post)
	sqlStr := `select 
    post_id,author_id,community_id,status,title,content,create_time
    from post 
    where post_id = ?`
	err = db.Get(post, sqlStr, pid)
	if err != nil {
		zap.L().Error("get post failed")
	}

	return
}
func GetPostList(p *models.ParamPostList) (posts []*models.Post, err error) {
	posts = make([]*models.Post, p.Size)
	page := (p.Page - 1) * p.Size
	size := p.Page * p.Size
	zap.L().Debug("", zap.Any("page", page), zap.Any("size", size))
	if p.Order == models.OrderByScore {
		sqlStr := `select 
    post_id,author_id,community_id,status,title,content,create_time,score
    from post 
    ORDER BY score DESC
    limit ?,?`
		err = db.Select(&posts, sqlStr, page, size)

		return
	} else {
		sqlStr := `select 
    post_id,author_id,community_id,status,title,content,create_time,score
    from post 
    ORDER BY create_time DESC 
    limit ?,?`
		err = db.Select(&posts, sqlStr, page, size)
		zap.L().Debug("", zap.Any("sda", posts))
		zap.L().Debug("", zap.Any("sda", posts))
		return
	}
}

func GetPostListByIds(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,title,content,author_id,community_id,create_time
    from post
    where post_id in (?)
    order by FIND_IN_SET(post_id,?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}

//func GetPostInOrder(p *models.ParamPostList) ([]string, error) {
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
