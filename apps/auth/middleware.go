package auth

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorp/core"
	"net/http"
	"strings"
	"time"
)

func LoginRequiredMiddleware(ctx *gin.Context) {
	tokenString, err := parseAuthHeader(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	claims, err := parseJwtToken(tokenString)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}
	exp, err := claims.GetExpirationTime()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": err.Error(),
		})
		return
	}

	if time.Now().Unix() > exp.Unix() {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Token Expired",
		})
	}

	user := User{}
	core.DB.First(&user, claims[core.Settings.UserObjId])
	if user.ID == 0 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid User",
		})
	}
	ctx.Set(core.Settings.UserObjKey, user)
	ctx.Next()
}

func parseAuthHeader(ctx *gin.Context) (string, error) {
	authHead := ctx.Request.Header.Get("Authorization")
	if authHead == "" {
		return "", errors.New("No valid header")
	}
	jwtArr := strings.Split(authHead, " ")
	if len(jwtArr) < 2 {
		return "", errors.New("No jwt provided")
	}
	token_string := jwtArr[1]
	return token_string, nil
}

func parseJwtToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(core.Settings.SecretKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("Unable to parse token")
}

func GetUser(ctx *gin.Context) (User, error) {
	user, err := ctx.Get("user")
	if err != true {
		return User{}, errors.New("Unable to get User")
	}
	return user.(User), nil
}
