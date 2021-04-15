package router

import (
	"context"
	"github.com/Peterliang233/debate/config"
	"github.com/Peterliang233/debate/middlerware"
	"github.com/Peterliang233/debate/router/v1/api/user"
	"github.com/Peterliang233/debate/router/v1/api/user/login"
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


	v1Group := router.Group("/v1")
	v1Group.POST("/login", login.Login)
	v1Group.POST("/registry", login.Registry)
	v1Group.POST("/verify", login.GetEmailCode)  //注册时点击获取邮箱验证码
	api := v1Group.Group("/user")
	api.Use(middlerware.JWTAuthMiddleware())
	{
		api.GET("/info", user.GetUser)         //查看用户信息
		api.PUT("/info", user.UpdateUser)   //修改用户信息
		api.PUT("/pwd", user.UpdatePassword) //修改用户的密码
		api.DELETE("/", login.Logout)    //登出
		api.POST("/upload", user.UpdatePhoto)  //上传图片
	}
	return router
}
