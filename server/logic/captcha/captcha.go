package captcha

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/tools"
)

func GetCaptcha(ctx *gin.Context) {
	captcha, err := tools.CaptchaGenerate()
	if err != nil {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "验证码生成失败!"})
		return
	}
	ctx.JSON(http.StatusOK, tools.ECode{Data: captcha})
}

func GetVerify(ctx *gin.Context) {
	var param tools.CaptchaData
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "参数绑定失败!"})
		return
	}
	if !tools.CaptchaVerify(param) {
		ctx.JSON(http.StatusOK, tools.ECode{Code: 1, Message: "验证失败!"})
		return
	}
	ctx.JSON(http.StatusOK, tools.ECode{Code: 0, Message: "ok!"})
}
