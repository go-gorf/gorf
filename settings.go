package main

import (
	"github.com/gin-gonic/gin"
	"gorp/core"
)
import "gorp/apps/auth"

// add all the apps
var apps = []core.GorpApp{
	&auth.AuthApp,
}

func LoadSettings() {
	// jwt secret key
	core.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	core.Apps = append(apps)
	LoadSettings()
	dbConf := core.DbConf{
		"data.db",
	}
	core.InitializeDatabase(&dbConf)
	core.SetupApps()
	r := gin.Default()
	core.RegisterApps(r)
	return r
}
