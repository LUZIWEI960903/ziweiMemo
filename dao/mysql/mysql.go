package mysql

import (
	"fmt"
	"ziweiMemo/settings"

	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	// root:123456@tcp(127.0.0.1:3306)/ziweiMemo?charset=utf8mb4&parseTime=True
	dsn := fmt.Sprintf("%s:%s@tcp(%s%s)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)

	// 连接mysql
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("[package: mysql] [func: Init] [sqlx.Connect(\"mysql\", dsn)] failed, err: %v\n", zap.Error(err))
		return
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return
}

func Close() {
	_ = db.Close()
}
