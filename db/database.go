package db

import (
	"fmt"
	"go-authapi-adv/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type DatabaseConnection struct {
	DB *gorm.DB
}

func (d *DatabaseConnection) ConnectDataBase() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("loading .env failed")
	}

	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	d.DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("connection error:", err)
	} else {
		fmt.Println("connected to database ", Dbdriver)
	}

	d.DB.AutoMigrate(&models.RegisterInput{})

}
func (d *DatabaseConnection) SaveUser(user models.User) (*models.User, error) {
	var err error

	err = d.DB.Create(&user).Error
	if err != nil {
		return &models.User{}, err
	}
	return &user, nil
}
