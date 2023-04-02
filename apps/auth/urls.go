package auth

import "github.com/gin-gonic/gin"

func Urls(r *gin.Engine) {
	r.POST("/signup", UserSignUp)
	r.POST("/login", UserLogin)
	r.GET("/profile", LoginRequiredMiddleware, UserProfile)
}
