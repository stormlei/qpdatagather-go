package db

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"qpdatagather/entity"
	"time"
)

var db *gorm.DB

func Init(conf *viper.Viper) {
	logLevel := logger.Info
	if gin.Mode() == gin.ReleaseMode {
		logLevel = logger.Error
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logLevel,    // 日志级别
			Colorful:      true,        // 禁用彩色打印
		},
	)

	var err error
	db, err = gorm.Open(mysql.Open(conf.GetString("db.conn")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect to database")
	}

	// 迁移 schema
	db.AutoMigrate(&entity.BleMac{})
}

func GetDb() *gorm.DB {
	return db
}
