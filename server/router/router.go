package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"server/logic/captcha"
	"server/logic/login"
	"server/logic/register"
	"server/logic/webSocket"
	"syscall"
	"time"
)

func NewRouter() {

	r := gin.Default()

	r.GET("/ws", webSocket.WcDemoV0)
	r.GET("/login", login.GetLogin)
	r.GET("/registration", registration.GetRegistration)
	r.GET("/logout", login.Logout)

	//账号密码验证码登录
	r.POST("/login/nap", login.NameAndPwdLogin)
	//发送验证码接口(手机号登录和绑定都要先调这个接口)
	r.POST("/sendCode", login.SendCode)
	//手机号验证码登录
	r.POST("/login/sms", login.SmsLogin)
	//微信扫码登录
	r.POST("/login/wx", login.WxLogin)

	//用户注册
	r.POST("/registration/register", registration.Registration)
	//手机号绑定
	r.POST("/registration/bindPhone", registration.BindPhone)

	{ //验证码接口
		r.GET("/captcha", captcha.GetCaptcha)
		r.POST("/captcha/verify", captcha.GetVerify)
	}

	server := &http.Server{Addr: ":8088", Handler: r}

	// 启动HTTP服务器
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 监听系统信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, os.Kill, syscall.SIGTERM)

	// 等待接收到信号
	<-quit
	log.Println("Shutdown server...")

	// 优雅关闭HTTP服务器
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
