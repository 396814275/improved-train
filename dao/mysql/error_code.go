package mysql

import "errors"

var (
	ErrUserExist       = errors.New("用户已存在")
	ErrUserNotExist    = errors.New("用户不存在")
	ErrInvalidPassword = errors.New("密码错误")
	ErrInvalidId       = errors.New("无效的ID")
)
