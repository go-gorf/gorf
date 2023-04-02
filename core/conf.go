package core

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gorp App interface
type GorpApp interface {
	Setup()
	Register(r *gin.Engine)
}

// array of all registerd apps
var Apps []GorpApp

// Setup all apps
func SetupApps() {
	fmt.Println("Configuring apps")
	for _, app := range Apps {
		println(app)
		app.Setup()
	}
}

// Register apps
func RegisterApps(r *gin.Engine) {
	fmt.Println("Registering apps")
	for _, app := range Apps {
		app.Register(r)
	}
}

// Global Project settings
type GlobalSettings struct {
	SecretKey  string
	UserObjKey string
	UserObjId  string
}

var Settings = GlobalSettings{
	UserObjKey: "user",
	UserObjId:  "id",
}
