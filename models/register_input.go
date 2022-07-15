package models

type RegisterInput struct {
	Mail     string `json:"mail" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Lastname string `json:"lastname" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
