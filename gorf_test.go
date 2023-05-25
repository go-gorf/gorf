package gorf

import (
	"encoding/json"
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

}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	Apps = append(apps)
	LoadSettings()
	_ = InitializeDatabase()

	SetupApps()
	r := gin.Default()
	RegisterApps(r)
	return r
}

func TestHealth(t *testing.T) {
	r := BootstrapRouter()
	req, _ := http.NewRequest("GET", "/health", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var result map[string]string

	err := json.Unmarshal(w.Body.Bytes(), &result)
	if err != nil {
		assert.Fail(t, "Unable to parse the response")
	}

	assert.Equal(t, result["status"], "ok")
}
