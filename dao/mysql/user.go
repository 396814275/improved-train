package mysql

import (
	"crypto/md5"
	"encoding/hex"
	"web2/models"
)

// 把每一步数据库操作封装成函数
// 等待logic层根据业务需求调用
const secret = "zazahui"

// 判断用户是否存在
func CheckUserExist(username string) (err error) {
	sqlstr := "select count(user_id) from user where username=?"
	var count int
	db.Get(&count, sqlstr, username)
	if count > 0 {
		return ErrUserExist
	}
	return
}

// 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	user.Password = encrypPassword(user.Password)
	sqlstr := "insert into user (user_id,username,password) values(?,?,?)"
	_, err = db.Exec(sqlstr, user.UserId, user.Username, user.Password)
	return
	//	执行SQL语句入库

}

func encrypPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func Login(user *models.User) (err error) {
	opassword := user.Password
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err != nil {
		//查询数据库失败
		return ErrUserNotExist
	}
	//    判断密码是否正确
	password := encrypPassword(opassword)
	if password != user.Password {
		return ErrInvalidPassword
	}
	return
}
func GetUserByAuthorID(uid int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select username from user where user_id=?`
	err = db.Get(user, sqlStr, uid)
	return
}
