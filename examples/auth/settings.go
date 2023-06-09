package main

import (
	"github.com/caarlos0/env/v8"
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/auth"
	"github.com/go-gorf/gorf"
	"github.com/go-gorf/gorf/backends/gormi"
	"log"
)

var apps = []gorf.App{
	&auth.App,
}

type config struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"GOo8Rs8ht7qdxvfwefnsn08YleBx6pEv0cQ09a"`
	Region    string `env:"REGION,required" envDefault:"eu-north-1"`
	ClientId  string `env:"CLIENT_ID,required"`
	UserPool  string `env:"USER_POOL,required"`
}

func LoadSettings() {
	cfg := &config{}
	err := env.Parse(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	//Gorf settings
	gorf.Settings.SecretKey = cfg.SecretKey
	gorf.Settings.DbBackends = gormi.NewSqliteBackend("db.sqlite")

	_ = gorf.InitializeDatabase()

	//Auth Settings
	auth.Settings.Region = cfg.Region
	auth.Settings.ClientId = cfg.ClientId
	auth.Settings.UserPool = cfg.UserPool
	auth.Settings.AuthMiddleware = auth.NewJwtMiddleware(
		gorf.DB,
		auth.NewCognitoJwks(),
	)
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
