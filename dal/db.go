package dal

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/rongpengju/gin-template/configs"
)

// DBWithDefault 当前项目的默认 DB 对象
// 如果需要操作其他数据源，比如需要操作用户数据库
// Example:
// var DBWithUser *gorm.DB
var DBWithDefault *gorm.DB

func init() {
	DBWithDefault = ConnectDB(configs.Conf.DataSource.MySQL.DsnWithDefault)
}

// InitGormGen 设置 gorm/gen 的 DB 对象
func InitGormGen() {
	// TODO 设置DB对象
	//query.SetDefault(DBWithDefault)
}

func ConnectDB(dsn string) *gorm.DB {
	db, err := gorm.Open(
		mysql.Open(dsn),
		&gorm.Config{
			Logger:         newGormLogger(),
			TranslateError: true,
		})
	if err != nil {
		panic(fmt.Errorf("链接数据库失败: %w，数据库DSN：%s", err, dsn))
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(fmt.Errorf("获取sqlDB失败：%w", err))
	}

	// 可打开的最大连接数
	sqlDB.SetMaxOpenConns(configs.Conf.DataSource.MySQL.MaxOpen)
	// 最大空闲连接数
	sqlDB.SetMaxIdleConns(configs.Conf.DataSource.MySQL.MaxIdle)
	// 连接空闲后在多长时间内可复用
	sqlDB.SetConnMaxLifetime(time.Duration(configs.Conf.DataSource.MySQL.MaxLifeTime) * time.Second)

	return db
}
