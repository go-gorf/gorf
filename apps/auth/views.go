package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorp/core"
	"net/http"
	"time"
)

func UserSignUp(ctx *gin.Context) {
	var body struct {
		Email    string `binding:"required"`
		Password string `binding:"required"`
	}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while creating password",
		})
	}
	newUser := User{
		Email:    body.Email,
		Password: string(passwordHash),
	}
	result := core.DB.Create(&newUser)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "New user created successfully",
		"user_id": newUser.ID,
		"email":   newUser.Email,
	})
}

func UserLogin(ctx *gin.Context) {
	var body struct {
		Email    string `binding:"required"`
		Password string `binding:"required"`
	}
	err := ctx.ShouldBindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	user := User{}
	core.DB.First(&user, "email = ?", body.Email)
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		core.Settings.UserObjKey: user.ID,
		"timestamp":              time.Now(),
		"exp":                    time.Now().Add(time.Hour * 24).Unix(),
	})

	token_string, err := token.SignedString([]byte(core.Settings.SecretKey))

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Unable to generate jwt token",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "JWT token generated successfully",
		"jwt":     token_string,
	})
}

func UserProfile(ctx *gin.Context) {
	user, _ := GetUser(ctx)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Success",
		"email":   user.Email,
	})
}
