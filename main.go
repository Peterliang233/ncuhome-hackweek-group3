package main

import (
	"github.com/Peterliang233/debate/router"
	"github.com/Peterliang233/debate/service/v1/dao"
)

func main() {
	//初始化数据库
	dao.InitDb()

	//运行服务
	router.Run()
}