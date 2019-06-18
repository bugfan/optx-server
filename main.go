package main

import (
	"log"
	"optx-server/apis"
	"optx-server/models"
	"optx-server/settings"
	"os"
)

func main() {
	// init database
	_, err := models.SetEngine(&models.Config{
		User:     settings.Get("db_user"),
		Password: settings.Get("db_password"),
		Host:     settings.Get("db_host"),
		Name:     settings.Get("db_name"),
		Log:      settings.Get("db_log"),
	})
	if err != nil {
		log.Fatal("链接数据库失败:", err)
		os.Exit(-1)
	}
	// run api server
	if settings.Get("ishttps") == "true" {
		apis.NewAPIServer().G.RunTLS(settings.Get("server_addr")+":"+settings.Get("server_port"), "./ssl.pem", "./ssl.key")
	} else {
		apis.NewAPIServer().G.Run(settings.Get("server_addr") + ":" + settings.Get("server_port"))
	}
}
