package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServeBusy
	CodeInvalidAuth
	CodeNeedLogin
)

var codeMsg = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNotExist:    "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServeBusy:       "服务繁忙",
	CodeNeedLogin:       "需要登录",
	CodeInvalidAuth:     "无效的的token",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsg[c]
	if !ok {
		msg = codeMsg[CodeInvalidParam]
	}
	return msg
}
