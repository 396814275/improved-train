package mysql

import "go.uber.org/zap"

func VoteForPost(userID string, postID string, value int, uid int64) (err error) {
	//1.接受参数并进行校验
	sqlStr := `select count(uid) 
			   from vote 
			   where uid = ? and post_id = ?`
	isExist := 0
	ov := 0
	err = db.Get(&isExist, sqlStr, uid, postID)
	//限制每个用户投一票
	if isExist == 0 {
		sqlStr = `insert into vote(
    uid,user_id,post_id,status)
    values (?,?,?,?) `
		_, err = db.Exec(sqlStr, uid, userID, postID, value)
	} else {
		sqlStr1 := `select status from vote 
			 where uid = ?`
		err = db.Get(&ov, sqlStr1, uid)
	}

	zap.L().Info("", zap.Any("ov", ov), zap.Error(err))
	//计算两次操作的差值
	var dir float64
	if value > ov {
		dir = 1
	} else if value < ov {
		dir = -1
	} else if value == ov && value == 1 {
		dir = -1
	} else {
		dir = 1
	}
	zap.L().Info("", zap.Any("dir:", dir), zap.Error(err))
	zap.L().Info("", zap.Any("postID:", postID), zap.Error(err))
	//	2.更新
	if dir == 1 {
		sqlStr2 := `update post set score=score+1
            where post_id=?`
		_, err = db.Exec(sqlStr2, postID)
	} else {
		sqlStr2 := `update post set score=score-1
            where post_id=?`
		_, err = db.Exec(sqlStr2, postID)
	}

	if err != nil {
		zap.L().Error("update user's vote failed", zap.Error(err))
	}
	return
}
