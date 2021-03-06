package router

import (
	"context"
	"github.com/Peterliang233/debate/config"
	"github.com/Peterliang233/debate/middlerware"
	debate "github.com/Peterliang233/debate/router/v1/api/debate"
	"github.com/Peterliang233/debate/router/v1/api/user"
	"github.com/Peterliang233/debate/router/v1/api/user/login"
	"github.com/Peterliang233/debate/service/v1/api/socket"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)


func Run() {
	r := InitRouter()
	readTimeout := config.ServerSetting.ReadTimeout
	writeTimeout := config.ServerSetting.WriteTimeout
	maxHeaderBytes := 1<<20
	server := &http.Server{
		Addr:           config.ServerSetting.HttpPort,
		Handler:        r,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}


	//优雅地关机和重启
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit

	log.Println("Shutdown Server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalln("server shutdown", err)
	}
	log.Println("Server exited")
}



func InitRouter() *gin.Engine{
	router := gin.New()
	gin.SetMode(config.ServerSetting.RunMode)
	router.Use(middlerware.Logger())
	router.Use(middlerware.Cors())

	v1Group := router.Group("/v1/api")
	v1Group.POST("/login", login.Login)
	v1Group.POST("/registry", login.Registry)
	v1Group.POST("/verify", login.GetEmailCode)  //注册时点击获取邮箱验证码
	//api := v1Group.Group("/api")
	hub := socket.NewHub()
	go hub.Run()
	v1Group.Use(middlerware.JWTAuthMiddleware())
	{
		//用户组
		V1User := v1Group.Group("/user")
		{
			V1User.GET("/info", user.GetUser)        //查看用户信息
			V1User.POST("/upload", user.UpdatePhoto) //上传图片
			V1User.PUT("/info", user.UpdateUser)     //修改用户信息
			V1User.PUT("/pwd", user.UpdatePassword)  //修改用户的密码
			//V1User.DELETE("/logout", login.Logout)         //登出
		}
		//辩论组
		V1Debate := v1Group.Group("/debate")
		{
			V1Debate.POST("/one", debate.OneToOneDebate)  //发送辩论记录
			V1Debate.GET("/record/:id", debate.GetRecord) //获取单个辩论记录
			V1Debate.POST("/records", debate.GetRecords)    //获取所有辩论记录
			//V1Debate.GET("/records/last")  //获取完成的辩论场
			//V1Debate.GET("/records/future")  //获取将来的辩论场
			V1Debate.POST("/pos", debate.ChosePositive)   //选择正方
			V1Debate.POST("/neg", debate.ChoseNegative)   //选择反方
			V1Debate.POST("/add", debate.AddTitle)  //添加辩题
		}
		//通信组
		router.GET("/ws", func(c *gin.Context){  //socket通信接口
			socket.ServeWs(hub, c)
		})
	}

	return router
}
