package gormi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-gorf/gorf"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var apps = []gorf.GorfApp{}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
	gorf.Settings.DbBackends = &GormSqliteBackend{}

}

// bootstrap server
func BootstrapRouter() *gin.Engine {
	gorf.Apps = append(apps)
	LoadSettings()
	_ = gorf.InitializeDatabase()

	gorf.SetupApps()
	r := gin.Default()
	gorf.RegisterApps(r)
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
