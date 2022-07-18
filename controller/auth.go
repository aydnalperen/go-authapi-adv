package controller

import (
	"go-authapi-adv/models"
	"html"
	"net/http"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User

	user.Email = input.Mail
	user.Name = input.Name
	user.Lastname = input.Lastname
	//user.Password = input.Password
	user.Username = input.Username

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	user.Password = string(hashedPassword) // user creation is done
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	_, err = user.SaveUser()

	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func Login(ctx *gin.Context) {

}

func Logout(ctx *gin.Context) {

}
