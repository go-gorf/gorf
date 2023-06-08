package gorf

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// App Gorp App interface
type App interface {
	Setup() error
	Register(r *gin.Engine)
}

// Apps array of all registerd apps
var Apps []App

// SetupApps Setup all apps
func SetupApps() {
	fmt.Println("Configuring apps")
	for _, app := range Apps {
		err := app.Setup()
		if err != nil {
			panic("Unable to configure apps")
		}
	}
}

// RegisterApps Register apps
func RegisterApps(r *gin.Engine) {
	registerInternalUrls(r)
	fmt.Println("Registering apps")
	for _, app := range Apps {
		app.Register(r)
	}
}

type BaseApp struct {
	Name         string
	RouteHandler func(r *gin.Engine)
	SetUpHandler func() error
}

func (app *BaseApp) Setup() error {
	fmt.Printf("Configuring the %v app", app.Name)
	err := app.SetUpHandler()
	if err != nil {
		return err
	}
	return nil
}

func (app *BaseApp) Register(r *gin.Engine) {
	app.RouteHandler(r)
}

// GlobalSettings Global Project settings
type GlobalSettings struct {
	SecretKey  string
	UserObjKey string
	UserObjId  string
	DbBackends DbBackend
}

var Settings = &GlobalSettings{
	UserObjKey: "user",
	UserObjId:  "id",
	DbBackends: nil,
}
