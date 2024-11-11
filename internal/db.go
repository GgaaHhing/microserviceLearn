package internal

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"testProject/test/microservice/model"
	"time"
)

var DB *gorm.DB
var err error

func InitDB() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		},
	)
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		"root", "root", "127.0.0.1", 3308, "DBTest")
	DB, err = gorm.Open(mysql.Open(conn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接成功！")
	err = DB.AutoMigrate(&model.Account{})
	if err != nil {
		fmt.Println("AutoMigrate失败：", err)
	}
}
