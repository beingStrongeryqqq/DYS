package mysql

import "errors"

var (
	ErrorUserExist       = errors.New("用户已存在")
	ErrorUserNotExist    = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("用户名或密码错误")
	ErrorNotAlive        = errors.New("用户未登录")
	ErrorHadLogin        = errors.New("用户已在其他设备登录")
	ErrorInvalidID       = errors.New("无效的id")
)
