package models

import (
	"errors"
	"go-authapi-adv/utils"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Email      string `gorm:"size:255;not null;unique" json:"email"`
	Name       string `gorm:"size:255;not null;unique" json:"name"`
	Lastname   string `gorm:"size:255;not null;unique" json:"lastname"`
	Username   string `gorm:"size:255;not null;unique" json:"username"`
	Password   string `gorm:"size:255;not null;unique" json:"password"`
	IsVerified int    `gorm:"size:255;not null;unique; default:0" json:"is_valid"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
func (u *User) SaveUser() (*User, error) {
	if result := DB.Create(&u); result.Error != nil {
		return &User{}, result.Error
	}
	return u, nil
}

func LoginCheck(username string, password string) (string, error) {
	var err error

	user := User{}

	err = DB.Model(User{}).Where("username=?", username).Take(&user).Error

	if err != nil {
		return "", err
	}
	err = VerifyPassword(password, user.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := utils.GenerateToken(user.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
func GetUserById(id uint) (User, error) {
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		return user, errors.New("User not found!")
	}

	user.Password = ""

	return user, nil
}
