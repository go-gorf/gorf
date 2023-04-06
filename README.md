# Golang Rest Framework

[![Go](https://github.com/go-gorf/gorf/actions/workflows/go.yml/badge.svg)](https://github.com/go-gorf/gorf/actions/workflows/go.yml)

Django inspired Golang Rest Framework

## Installation
```shell
go get github.com/go-gorf/gorf
```
Install gorf auth app
```shell
go get github.com/go-gorf/auth
```

## main.go

Firstly, Create a new main package with following code

``` go title="main.go" 
package main

import (
	"log"
)

func main() {
	r := BootstrapRouter()
	err := r.Run()
	if err != nil {
		log.Fatal("Unable to create the gin server")
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```

## settings.go
Next, Create a settings.go file in the main package with the following code snippet

``` go title="settings.go"
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/auth"
	"github.com/go-gorf/gorf"
)

// add all the apps
var apps = []gorf.GorfApp{
	&auth.AuthApp,
}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
	gorf.Settings.DbConf = &gorf.SqliteBackend{
		Name: "db.sqlite",
	}
}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	gorf.Apps = append(apps)
	LoadSettings()
	gorf.InitializeDatabase()
	gorf.SetupApps()
	r := gin.Default()
	gorf.RegisterApps(r)
	return r
}
```
