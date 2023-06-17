package gorf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, resp gin.H) {
	ctx.JSON(http.StatusOK, resp)
}

func BadRequest(ctx *gin.Context, msg string, err error) {
	e := &Err{msg, err, http.StatusBadRequest}
	ctx.JSON(http.StatusBadRequest, e.Response())
}

func ErrorResponse(ctx *gin.Context, msg string, status int) {
	e := NewErr(msg, status, nil)
	ctx.JSON(status, e.Response())
}
