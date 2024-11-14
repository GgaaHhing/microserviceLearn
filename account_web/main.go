package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"testProject/microservice/account_web/handler"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "ip 输入")
	port := flag.Int("port", 8081, "port 输入")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	r := gin.Default()
	accountGroup := r.Group("/v1/account")
	{
		accountGroup.GET("/list", handler.AccountListHandler)
		accountGroup.POST("/login", handler.LoginByPasswordHandler)
		//让浏览器获取到我们发送的图片验证码
		accountGroup.GET("/captcha", handler.CaptchaHandler)
	}
	r.GET("/health", handler.HealthHandler)
	r.Run(addr)
}
