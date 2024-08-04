package controllers

import (
	"DYS/dao/mysql"
	"DYS/logic"
	"DYS/models"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func SignUpHandler(c *gin.Context) {
	//1.获取参数和参数校验
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("signup with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}

	fmt.Println(p)
	//2.业务处理
	err := logic.SignUp(p)
	if err != nil {
		zap.L().Error("signup failed, err:", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CodeUserExist)
			return
		}
		ResponseError(c, CodeServerBusy)
		return
	}
	//3.返回响应
	ResponseSuccess(c, nil)
}

func LoginHandler(c *gin.Context) {
	//获取参数及参数校验
	p := new(models.ParamLogin)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid param", zap.Error(err))
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(c, CodeInvalidParam, removeTopStruct(errs.Translate(trans)))
		return
	}
	//业务逻辑处理
	token, err, user := logic.Login(p)
	user.Token = token

	if err != nil {
		zap.L().Error("logic.login failed", zap.String("Username", p.Username), zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CodeUserNoExist)
			return
		}
		if errors.Is(err, mysql.ErrorHadLogin) {
			ResponseError(c, CodeHadLogin)
			return
		}
		ResponseError(c, CodeInvalidPassword)
		return
	}
	//返回响应
	ResponseSuccess(c, gin.H{
		"user_id":   fmt.Sprintf("%d", user.UserID), // id值大于1<<53-1  int64类型的最大值是1<<63-1
		"user_name": user.Username,
		"token":     user.Token,
	})
}
