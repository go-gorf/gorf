package gorf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Err struct {
	Msg    string
	Er     error
	status int
}

func (e *Err) Response() gin.H {
	return gin.H{
		"Message":    e.Msg,
		"Error":      e.Er.Error(),
		"StatusCode": e.status,
	}
}

func (e *Err) Error() string {
	return e.Er.Error()
}

func BadRequest(ctx *gin.Context, msg string, err error) {
	e := &Err{msg, err, http.StatusBadRequest}
	ctx.JSON(http.StatusBadRequest, e.Response())
}