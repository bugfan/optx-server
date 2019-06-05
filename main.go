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
	apis.NewAPIServer().G.Run(":9997")
}
