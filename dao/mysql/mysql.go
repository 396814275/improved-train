package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
	"web2/settings"
)

var db *sqlx.DB

func Init(cfg *settings.Mysqlconfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		//viper.GetString("mysql.user"),
		//viper.GetString("mysql.password"),
		//viper.GetString("mysql.host"),
		//viper.GetString("mysql.port"),
		//viper.GetString("mysql.dbname"),
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Dbname,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect to DB failed, err:%v\n", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(settings.Conf.MaxOpenConns)
	db.SetMaxIdleConns(settings.Conf.MaxIdleConns)
	return
}
func Close() {
	if err := db.Close(); err != nil {
		fmt.Printf("close DB failed, err:%v\n", zap.Error(err))
	}
}
