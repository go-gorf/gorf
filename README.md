# Golang Rest Framework

[![Go](https://github.com/go-gorf/gorf/actions/workflows/go.yml/badge.svg)](https://github.com/go-gorf/gorf/actions/workflows/go.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-gorf/gorf.svg)](https://pkg.go.dev/github.com/go-gorf/gorf)  
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

``` go
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

## Write your own apps

Create a new package named "hello"

add app.go file

``` go
package hello

import (
	"github.com/go-gorf/gorf"
)

func setup() error {
	// Add setup here
	return nil
}

var HelloApp = gorf.GorfBaseApp{
	Name:         "Hello",
	RouteHandler: Urls,
	SetUpHandler: setup,
}

```

add urls.go file

``` go
package hello

import "github.com/gin-gonic/gin"

func Urls(r *gin.Engine) {
	r.POST("/hello", Hello)
}
```

add views.go file

``` go
package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Hello world",
	})
}

```
## Check the gorf app template repo

https://github.com/go-gorf/template

## Development

```bash
go mod edit -replace github.com/go-gorf/auth=../auth
```
Or

```
replace github.com/go-gorf/gorf => ../gorf
replace github.com/go-gorf/auth => ../auth
```