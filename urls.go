package gorf

import "github.com/gin-gonic/gin"

func registerInternalUrls(r *gin.Engine) {
	r.GET("/health", Health)
}
