package main

import (
	"go-authapi-adv/controller"
	"go-authapi-adv/middleware"
	"go-authapi-adv/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.ConnectDataBase()
	r := gin.Default()

	public := r.Group("/auth")

	public.POST("/login", controller.Login)
	public.POST("/register", controller.Register)

	protected := r.Group("/admin")
	protected.Use(middleware.JwtAuthMiddleWare())
	protected.GET("/user", controller.CurrentUser)
	protected.POST("/logout", controller.Logout)

	r.Run(":8080")
}
