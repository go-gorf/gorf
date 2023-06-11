package gorf

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Response(ctx *gin.Context, resp gin.H) {
	ctx.JSON(http.StatusOK, resp)
}
