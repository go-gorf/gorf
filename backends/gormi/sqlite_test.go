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

var backend *GormSqliteBackend

func getBackend() *GormSqliteBackend {
	if backend != nil {
		return backend
	}
	backend := &GormSqliteBackend{"db.sqlite"}
	return backend

}

func LoadSettings() {
	// jwt secret key
	gorf.Settings.SecretKey = "GOo8Rs8ht7qdxv6uUAjkQuopRGnql2zWJu08YleBx6pEv0cQ09a"
	gorf.Settings.DbBackends = &GormSqliteBackend{"db.sqlite"}

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

func TestGorfHealth(t *testing.T) {
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

func TestGormSqliteBackend_Connect(t *testing.T) {
	backend := getBackend()
	_, err := backend.Connect()
	assert.NoError(t, err)
}

func TestGormSqliteBackend_Disconnect(t *testing.T) {
	backend := getBackend()
	err := backend.Disconnect()
	assert.NoError(t, err)
}

func TestGormDB_AutoMigrate(t *testing.T) {
	backend := getBackend()
	db, err := backend.Connect()
	assert.NoError(t, err)
	user := &User{}
	err = db.AutoMigrate(user)
	assert.NoError(t, err)
}

func TestGormDB_Create(t *testing.T) {
	backend := getBackend()
	db, err := backend.Connect()
	assert.NoError(t, err)
	user := &User{}
	err = db.AutoMigrate(user)
	assert.NoError(t, err)
	user.FirstName = "Go"
	user.LastName = "Gorf"
	err = db.Create(&user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)
}

func TestGormDB_First(t *testing.T) {
	backend := getBackend()
	db, err := backend.Connect()
	assert.NoError(t, err)
	user := &User{}
	err = db.AutoMigrate(user)

	assert.NoError(t, err)
	user.FirstName = "Go"
	user.LastName = "Gorf"

	err = db.Create(&user)
	assert.NoError(t, err)
	assert.NotZero(t, user.ID)

	newUser := &User{}
	db.First(newUser, user.ID)
	assert.Equal(t, newUser.FirstName, "Go")
}
