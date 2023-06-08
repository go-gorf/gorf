package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/admin"
	"github.com/go-gorf/gorf"
)

var apps = []gorf.App{
	&admin.App,
}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRvjndsvnsn08YleBx6pEv0cQ09a"

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
