# gorf
[![Go](https://github.com/go-gorf/gorf/actions/workflows/go.yml/badge.svg)](https://github.com/go-gorf/gorf/actions/workflows/go.yml)

Django inspired Golang Rest Framework

## Installation

Install gorf
```shell
github.com/go-gorf/gorf
```
Install gorf auth app
```shell
github.com/go-gorf/gorf
```

## Quickstart

Firstly, Create a new main package with following code

``` go
package main

import (
	"log"

	"github.com/go-gorf/gorf"
	"github.com/go-gorf/gorf-contrib/auth"
)

func main() {
	r := BootstrapRouter()
	user := auth.User{}
	println(user.Email)
	err := r.Run()
	if err != nil {
		log.Fatal("Unable to create the gin server")
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

```

Next, Create a settings.go file in the main package with the following code snippet

``` go
package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/gorf"
	"github.com/go-gorf/gorf-contrib/auth"
)

// add all the apps
var apps = []gorf.GorfApp{
	&auth.AuthApp,
}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	gorf.Apps = append(apps)
	LoadSettings()
	dbConf := gorf.DbConf{
		Name: "data.db",
	}
	gorf.InitializeDatabase(&dbConf)
	gorf.SetupApps()
	r := gin.Default()
	gorf.RegisterApps(r)
	return r
}
```

## Project Management Board

https://github.com/orgs/go-gorf/projects/1