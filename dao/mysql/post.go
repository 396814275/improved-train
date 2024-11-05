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
func GetPostList(page, size int64) (posts []*models.Post, err error) {

	sqlStr := `select 
    post_id,author_id,community_id,status,title,content,create_time
    from post 
    ORDER BY create_time DESC
    limit ?,?`
	posts = make([]*models.Post, 0, 3)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
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
