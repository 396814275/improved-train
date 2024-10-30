package logic

import (
	"web2/dao/mysql"
	"web2/models"
	"web2/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//	1.生成postid
	p.ID = int64(snowflake.GenID())
	//	保存到数据库
	return mysql.CreatePost(p)

}
