package model

import (
	"fmt"
	"github.com/Peterliang233/debate/config"
	"github.com/Peterliang233/debate/dao"
	"github.com/Peterliang233/debate/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

func InitDb() {
	var err error
	dao.Db, err = gorm.Open(config.DatabaseSetting.Type,
		fmt.Sprintf( "%s:%s@tcp(%s:%s)/%s?charset=utf8&paeseTime=True&loc=Local",
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
	dao.Db.AutoMigrate(&model.Identify{})

	dao.Db.DB().SetMaxIdleConns(10)
	dao.Db.DB().SetMaxOpenConns(100)
	dao.Db.DB().SetConnMaxLifetime(10 * time.Second)
}