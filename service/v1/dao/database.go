package dao

import (
	"fmt"
	"github.com/Peterliang233/debate/config"
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/model"
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func InitDb() {
	InitRedis()  //初始化redis数据库
	InitMysql()  //初始化mysql数据库
}

func InitMysql() {
	var err error
	dao.Db, err = gorm.Open(config.DatabaseSetting.Type,
		fmt.Sprintf( "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			config.DatabaseSetting.User,
			config.DatabaseSetting.Password,
			config.DatabaseSetting.Host,
			config.DatabaseSetting.Port,
			config.DatabaseSetting.Dbname,
		))
	if err != nil {
		log.Fatalf("数据库打开失败")
	}

	dao.Db.SingularTable(true)
	dao.Db.AutoMigrate(&model.User{})

	dao.Db.DB().SetMaxIdleConns(10)
	dao.Db.DB().SetMaxOpenConns(100)
	dao.Db.DB().SetConnMaxLifetime(10 * time.Second)
}

func InitRedis() {
	var err error
	dao.Conn, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		log.Println("connect redis error :", err)
	}
}