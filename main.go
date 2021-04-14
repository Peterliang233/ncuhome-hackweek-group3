package main

import (
	"github.com/Peterliang233/debate/router"
	"github.com/Peterliang233/debate/service/v1/model"
)

func main() {
	//初始化数据库
	model.InitDb()

	//运行服务
	router.Run()
}