package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/auth"
	"github.com/go-gorf/gorf"
)

var apps = []gorf.App{
	&auth.App,
}

func LoadSettings() {
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxvfwefnsn08YleBx6pEv0cQ09a"
}

// BootstrapRouter bootstrap server
func BootstrapRouter() *gin.Engine {
	gorf.Apps = append(apps)
	LoadSettings()
	_ = gorf.InitializeDatabase()

	gorf.SetupApps()
	r := gin.Default()
	gorf.RegisterApps(r)
	return r
}
