package login

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/db"
	"server/model"
	"server/tools"
)

func GetLogin(context *gin.Context) {
	context.HTML(http.StatusOK, "login.html", nil)
}

type loginUser struct {
	Name         string `json:"name"`
	Password     string `json:"password"`
	CaptchaId    string `json:"captcha_id"`
	CaptchaValue string `json:"captcha_value"`
}

func NameAndPwdLogin(ctx *gin.Context) {
	var lUser loginUser
	if err := ctx.ShouldBind(&lUser); err != nil {
		ctx.JSON(http.StatusBadRequest, tools.ECode{Code: 1, Message: err.Error()})
		return
	}
	pwd := tools.Encrypt(lUser.Password)
	user := model.GetUserByNameAndPassword(lUser.Name, pwd)
	if user == nil || user.Password != pwd || user.Name != lUser.Name {
		ctx.JSON(http.StatusUnauthorized, tools.ECode{Code: 1, Message: "账号或密码错误！"})
		return
	}
	//if !tools.CaptchaVerify(tools.CaptchaData{CaptchaId: lUser.CaptchaId, Data: lUser.CaptchaValue}) {
	//	ctx.JSON(http.StatusPreconditionFailed, tools.ECode{Code: 1, Message: "验证码校验失败！"})
	//	return
	//}
	//设置Session到Redis
	_ = db.SetSessionLogin(ctx, user.Name, user.Uid)
	ctx.JSON(http.StatusOK, tools.ECode{Code: 0, Message: "登录成功！"})
}

func Logout(ctx *gin.Context) {
	_ = db.FlushSessionLogin(ctx)
}
