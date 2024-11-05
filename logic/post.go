package logic

import (
	"go.uber.org/zap"
	"web2/dao/mysql"
	"web2/dao/redis"
	"web2/models"
	"web2/pkg/snowflake"
)

func CreatePost(p *models.Post) (err error) {
	//	1.生成postid
	p.ID = int64(snowflake.GenID())
	//	保存到数据库
	err = mysql.CreatePost(p)
	err = redis.CreatePost(p.ID)
	return
}

// GetPostByID 根据帖子id返回帖子详情数据
func GetPostByID(pid int64) (data *models.ApiPost, err error) {
	//查询并组合我们接口想用的数据
	data = new(models.ApiPost)
	post, err := mysql.GetPostByID(pid)
	if err != nil {
		zap.L().Error("mysql.GetPostByID(pid) failed", zap.Int64("pid", pid), zap.Error(err))
		return nil, err
	}
	//根据作者ID获得作者信息
	user, err := mysql.GetUserByAuthorID(post.AuthorID)
	if err != nil {
		zap.L().Error("mysql.GetUserByAuthorID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	//根据社区ID获取社区详细信息
	community, err := mysql.GetCommunityDetailByID(post.CommunityID)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
		return
	}
	data = &models.ApiPost{
		AuthorName:      user.Username,
		Post:            post,
		CommunityDetail: community,
	}

	return
}

// GetPostList 获取帖子列表
func GetPostList(page, size int64) (data []*models.ApiPost, err error) {
	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		zap.L().Error("mysql.GetPostList() failed", zap.Error(err))
		return nil, err
	}
	data = make([]*models.ApiPost, 0, len(posts))
	for _, post := range posts {
		user, err := mysql.GetUserByAuthorID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByAuthorID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		//根据社区ID获取社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		postdetail := &models.ApiPost{
			AuthorName:      user.Username,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postdetail)
	}

	return

}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPost, err error) {
	//	去redis查询id列表
	ids, err := redis.GetPostInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		zap.L().Warn("redis.GetPostIdsInOrder(p) return 0 data")
		return
	}
	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
	//根据id去MySQL数据库查询帖子信息
	//返回的数据还要按照我给定的顺序排序
	posts, err := mysql.GetPostListByIds(ids)
	if err != nil {
		return
	}
	//提前查询好每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}
	//将帖子的作者和分区信息查询出来并返回
	for idx, post := range posts {
		user, err := mysql.GetUserByAuthorID(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.GetUserByAuthorID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		//根据社区ID获取社区详细信息
		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
			continue
		}
		postdetail := &models.ApiPost{
			AuthorName:      user.Username,
			VoteNum:         voteData[idx],
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postdetail)
	}
	return
}

//func GetCommunityPostList2(p *models.ParamPostList) (data []*models.ApiPost, err error) {
//	//	去redis查询id列表
//	ids, err := redis.GetCommunityPostIDsInOrder(p)
//	if err != nil {
//		return
//	}
//	if len(ids) == 0 {
//		zap.L().Warn("redis.GetPostIdsInOrder(p) return 0 data")
//		return
//	}
//	zap.L().Debug("GetPostList2", zap.Any("ids", ids))
//	//根据id去MySQL数据库查询帖子信息
//	//返回的数据还要按照我给定的顺序排序
//	posts, err := mysql.GetPostListByIds(ids)
//	if err != nil {
//		return
//	}
//	//提前查询好每篇帖子的投票数
//	voteData, err := redis.GetPostVoteData(ids)
//	if err != nil {
//		return
//	}
//	//将帖子的作者和分区信息查询出来并返回
//	for idx, post := range posts {
//		user, err := mysql.GetUserByAuthorID(post.AuthorID)
//		if err != nil {
//			zap.L().Error("mysql.GetUserByAuthorID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
//			continue
//		}
//		//根据社区ID获取社区详细信息
//		community, err := mysql.GetCommunityDetailByID(post.CommunityID)
//		if err != nil {
//			zap.L().Error("mysql.GetCommunityDetailByID(post.AuthorID) failed", zap.Int64("author_id", post.AuthorID), zap.Error(err))
//			continue
//		}
//		postdetail := &models.ApiPost{
//			AuthorName:      user.Username,
//			VoteNum:         voteData[idx],
//			Post:            post,
//			CommunityDetail: community,
//		}
//		data = append(data, postdetail)
//	}
//	return
//}
