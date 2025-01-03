package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"northal.com/api"
	"northal.com/config"
	_ "northal.com/docs"
	"northal.com/internal/biz"
	"northal.com/internal/server"
)

//	@title			权限管理系统
//	@version		1.0
//	@description	权限管理系统api、测试
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	North
//	@contact.email	1227379879@qq.com

//	@servers					[{"url":"http://localhost:8080", "description":"Local development server"}]
//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
//	@description				通过在请求头中添加 Authorization 字段进行认证

func main() {
	config.InitViper()
	router := server.InitHttp()
	db := server.InitDatabase()
	redis := server.InitRedis()

	db.AutoMigrate(&biz.User{}, &biz.Role{}, &biz.Menu{}, &biz.UserRole{})

	// 加载所有api
	api.NewSetupApi(router, db, redis).SetupRoutes()

	// 加载swagger
	server.InitDocs(router)

	srv := &http.Server{
		Addr:    ":" + config.GetAppConfig().Port,
		Handler: router,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
