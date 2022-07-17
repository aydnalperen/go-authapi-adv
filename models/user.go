package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Name     string `gorm:"size:255;not null;unique" json:"name"`
	Lastname string `gorm:"size:255;not null;unique" json:"lastname"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;unique" json:"password"`
}
