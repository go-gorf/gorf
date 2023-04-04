package gorf

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var apps = []GorfApp{}

func LoadSettings() {
	// jwt secret key
	Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
	Settings.DbConf = &SqliteBackend{
		"db.sqlite",
	}
}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	Apps = append(apps)
	LoadSettings()
	InitializeDatabase()
	SetupApps()
	r := gin.Default()
	RegisterApps(r)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "ok"})
	})
	return r
}

func TestNewUserHandler(t *testing.T) {
	r := BootstrapRouter()
	req, _ := http.NewRequest("GET", "/test", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
