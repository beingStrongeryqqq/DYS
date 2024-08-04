package mysql

import (
	"DYS/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
)

const secret = "yq"

// CheckUserExist 把每一步数据库操作封装成函数，等待logic层调用
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err = db.Get(&count, sqlStr, username)
	if err != nil {
		return
	}
	if count > 0 {
		return ErrorUserExist
	}

	return
}

func IsFirstLogin(token *models.Token) bool {
	sqlStr := "select count(*) from token where user_id = ?"
	var count int
	err := db.Get(&count, sqlStr, token.UserID)
	if err != nil {
		return false
	}
	if count < 1 {
		return true
	}

	return false
}

func InsertToken(token *models.Token) (err error) {
	sqlStr := `insert into token(user_id, tokendata) values (?,?)`
	_, err = db.Exec(sqlStr, token.UserID, token.TokenData)
	return
}

func DeleteToken(token *models.Token) (err error) {
	sqlStr := "DELETE FROM token WHERE user_id = ?"
	_, err = db.Exec(sqlStr, token.UserID)
	return
}

func InsertUser(user *models.User) (err error) {
	//对密码进行加密 mysql密码不支持明文
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id, username, password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return
}

func encryptPassword(op string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(op)))
}

func GetUserID(username string) (userid int64, err error) {
	sqlStr := `select user_id from user where username = ?`
	var userID int64
	err = db.Get(&userID, sqlStr, username)
	if err == sql.ErrNoRows {
		return 0, ErrorNotAlive
	}
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func CheckToken(token *models.Token) (err error) {
	oT := token.TokenData
	sqlStr1 := `select user_id from token where tokendata = ?`
	err = db.Get(token, sqlStr1, token.TokenData)
	if err == sql.ErrNoRows {
		return ErrorNotAlive
	}

	if err != nil {
		return err
	}
	if oT != token.TokenData {
		return ErrorHadLogin
	}

	return nil
}

func Login(user *models.User, token *models.Token) error {
	oP := user.Password
	sqlStr := `select user_id, username, password from user where username = ?`
	err := db.Get(user, sqlStr, user.Username)

	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}

	if err != nil {
		return err
	}

	err = CheckToken(token)
	if err != nil {
		return err
	}

	password := encryptPassword(oP)
	if password != user.Password {
		return ErrorInvalidPassword
	}

	return err
}
