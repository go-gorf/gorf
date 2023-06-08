package gorf

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// Apps array of all registered apps
var Apps []App

// App Gorp App interface
type App interface {
	Name() string
	Description() string
	Setup() error
	Register(r *gin.Engine)
}

type BaseApp struct {
	Title        string
	Info         string
	RouteHandler func(r *gin.Engine)
	SetUpHandler func() error
}

func (app *BaseApp) Name() string {
	return app.Title
}
func (app *BaseApp) Description() string {
	return app.Info
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
