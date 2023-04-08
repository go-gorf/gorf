package gorf

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Gorp App interface
type GorfApp interface {
	Setup() error
	Register(r *gin.Engine)
}

// array of all registerd apps
var Apps []GorfApp

// Setup all apps
func SetupApps() {
	fmt.Println("Configuring apps")
	for _, app := range Apps {
		err := app.Setup()
		if err != nil {
			panic("Unable to configure apps")
		}
	}
}

// Register apps
func RegisterApps(r *gin.Engine) {
	fmt.Println("Registering apps")
	for _, app := range Apps {
		app.Register(r)
	}
}

type GorfBaseApp struct {
	Name         string
	RouteHandler func(r *gin.Engine)
	SetUpHandler func() error
}

func (app *GorfBaseApp) Setup() error {
	fmt.Printf("Configuring the %v app", app.Name)
	err := app.SetUpHandler()
	if err != nil {
		return err
	}
	return nil
}

func (app *GorfBaseApp) Register(r *gin.Engine) {
	app.RouteHandler(r)
}

// Global Project settings
type GlobalSettings struct {
	SecretKey  string
	UserObjKey string
	UserObjId  string
	DbConf     DatabaseBackend
}

var Settings = GlobalSettings{
	UserObjKey: "user",
	UserObjId:  "id",
	DbConf:     &SqliteBackend{Name: "data.db"},
}
