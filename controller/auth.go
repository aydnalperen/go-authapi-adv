package controller

import (
	"go-authapi-adv/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var input models.RegisterInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "validated!"})
}

func Login(ctx *gin.Context) {

}

func Logout(ctx *gin.Context) {

}
