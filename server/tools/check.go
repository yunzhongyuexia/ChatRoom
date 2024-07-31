package tools

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"server/db"
)

// CheckPhone 验证手机号码是否符合中国大陆的手机号码格式
func CheckPhone(phone string) bool {
	// 正则表达式匹配中国大陆手机号码
	// 1开头，第二位3-9，后面跟随9位数字
	reg := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return reg.MatchString(phone)
}

func CheckUser(ctx *gin.Context) {
	var uid int64
	var name string
	session := db.GetSessionLogin(ctx)
	if v, ok := session["name"]; ok {
		name = v.(string)
	}
	if v, ok := session["uid"]; ok {
		uid = v.(int64)
	}
	if name == "" || uid <= 0 {
		ctx.JSON(http.StatusPreconditionFailed, ECode{Code: 1, Message: "用户未登录！"})
		return
	}
	ctx.Next()
}
