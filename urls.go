package gorf

import "github.com/gin-gonic/gin"

func internalUrls(r *gin.Engine) {
	r.GET("/health", Health)
}
