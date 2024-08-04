package logic

import (
	"DYS/dao/mysql"
	"DYS/models"
	"DYS/pkg/jwt"
	"DYS/pkg/snowflake"
)

// SignUp 存放业务逻辑代码
func SignUp(p *models.ParamSignUp) (err error) {
	//判断用户存不存在
	err = mysql.CheckUserExist(p.Username)
	if err != nil {
		//数据库查询出错
		return
	}
	//生成UID
	userID := snowflake.GenID()
	//构造一个user实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//保存进数据库
	return mysql.InsertUser(user)
}

func SetToken(userid int64, tokendata string) *models.Token {
	return &models.Token{
		UserID:    userid,
		TokenData: tokendata,
	}
}

func Login(p *models.ParamLogin) (token string, err error, user *models.User) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}

	userid, _ := mysql.GetUserID(user.Username)

	token, err = jwt.GenToken(userid, user.Username)

	tokenpass := SetToken(userid, token)

	if mysql.IsFirstLogin(tokenpass) {
		_ = mysql.InsertToken(tokenpass)
	} else {
		_ = mysql.DeleteToken(tokenpass)
		return "", mysql.ErrorHadLogin, nil
	}

	if err = mysql.Login(user, tokenpass); err != nil {
		return "", err, nil
	}

	return

}
