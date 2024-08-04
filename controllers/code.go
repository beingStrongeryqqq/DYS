package controllers

type ResCode int64

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNoExist
	CodeInvalidPassword
	CodeServerBusy

	CodeInvalidToken
	CodeNeedLogin
	CodeHadLogin
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "成功",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户名已存在",
	CodeUserNoExist:     "用户名不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
	CodeNeedLogin:       "需要登陆",
	CodeInvalidToken:    "无效的token",
	CodeHadLogin:        "用户已在其他设备登录",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
