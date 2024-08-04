package mysql

import (
	"DYS/settings"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 不要忘了导入数据库驱动
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	//"user:password@tcp(127.0.0.1:3306)/dbname"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	// 也可以使用MustConnect连接不成功就panic
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
