package models

type User struct {
	UserID   int64  `db:"user_id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Token    string
}

type Token struct {
	UserID    int64  `db:"user_id"`
	TokenData string `db:"tokendata"`
}
