package db

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// session包自带的结构体
//type Session struct {
//		ID      string
//		Values  map[interface{}]interface{}
//		Options *Options
//		IsNew   bool
//		store   Store
//		name    string
//}

var sessionLogin = "session-login"
var sessionCode = "session-code"

func SetSessionLogin(ctx *gin.Context, name string, uid int64) error {
	store := NewRedisStoreLogin(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionLogin)
	session.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	session.Values["name"] = name
	session.Values["uid"] = uid
	return session.Save(ctx.Request, ctx.Writer)
}

func GetSessionLogin(ctx *gin.Context) map[interface{}]interface{} {
	store := NewRedisStoreLogin(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionLogin)
	return session.Values
}

func FlushSessionLogin(ctx *gin.Context) error {
	store := NewRedisStoreLogin(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionLogin)
	session.Values["name"] = ""
	session.Values["uid"] = int64(0)
	return session.Save(ctx.Request, ctx.Writer)
}

func SetSessionCode(ctx *gin.Context, phone string, code string) error {
	store := NewRedisStoreCode(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionCode)
	session.ID = strconv.FormatInt(time.Now().UnixNano(), 10)
	session.Values["phone"] = phone
	session.Values["code"] = code
	return session.Save(ctx.Request, ctx.Writer)
}

func GetSessionCode(ctx *gin.Context) map[interface{}]interface{} {
	store := NewRedisStoreCode(ctx, Redis)
	session, _ := store.Get(ctx.Request, sessionCode)
	return session.Values
}
