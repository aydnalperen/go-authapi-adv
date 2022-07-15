package main

import (
	"go-authapi-adv/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	public := r.Group("/auth")

	public.POST("/login", controller.Login)
	public.POST("/register", controller.Register)
	public.POST("/logout", controller.Logout)

	r.Run(":8080")
}
