# Golang Rest Framework (Gorf)

[![Build Status](https://img.shields.io/github/actions/workflow/status/go-gorf/gorf/go.yml?branch=main&style=for-the-badge)](https://github.com/go-gorf/gorf/actions?workflow=build)
[![](http://img.shields.io/badge/godoc-reference-5272B4.svg?style=for-the-badge)](https://pkg.go.dev/github.com/go-gorf/gorf) 

Django inspired Rest Framework for Golang.

## Vision

Introducing Gorf, the Golang REST API framework that empowers developers to easily create and reuse modular apps. 
Inspired by the simplicity and flexibility of Python Django, Gorf utilizes the full MVC architecture, providing
dedicated model, URL, and view files for effortless understanding. Gorf's plug and play app concept enables infinite 
scalability and customization, making Gorf the go-to solution for anyone looking to kickstart a new project in Golang.
Join our community and contribute your ideas and innovations to build a framework that streamlines the development process for all.

Features

* Generic
* Simplicity
* Based on known architecture
* Reusability
* Support for apps
* Custom database backends
* More feature work in progress

## Installation
```bash
go get github.com/go-gorf/gorf
```
Install gorf auth app
```bash
go get github.com/go-gorf/auth
```

## Tutorial

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
## Check the Gorf project template


[https://github.com/go-gorf/template](https://github.com/go-gorf/template)

