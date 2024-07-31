package main

import (
	"server/config"
	"server/db"
	"server/router"
)

func main() {
	config.NewViper()
	db.NewMysql()
	db.NewRedis()
	defer func() {
		_ = db.MClose()
		_ = db.RClose()
	}()
	router.NewRouter()
}
