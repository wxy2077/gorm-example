package initialize

import (
	"fmt"
	"gorm-example/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func GormMysql(config *config.MySQL) *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       fmt.Sprintf("%s/%s?charset=utf8mb4&parseTime=True&loc=Local", config.DNS, config.Database), // DSN data source name
		DefaultStringSize:         256,                                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                      // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})

	if err != nil {
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(config.MaxIDleConn)
		sqlDB.SetMaxOpenConns(config.MaxOpenConn)
		sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime) * time.Second)
		return db
	}
}
