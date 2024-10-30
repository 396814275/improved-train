package logic

import (
	"web2/dao/mysql"
	"web2/models"
)

func GetCommunityList() ([]*models.Community, error) {
	//	查找数据库中所有的community数据
	return mysql.GetCommunityList()
	//	将数据返回
}
func GetCommunityDetailList(id int64) (*models.CommunityDetail, error) {
	//	查找数据库中所有的communityDetail数据
	return mysql.GetCommunityListByID(id)
	//	将数据返回
}
