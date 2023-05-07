package userDB

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"unique;not null"`
	Password     string `gorm:"not null"`
	IsAdmin      bool   `gorm:"not null"`
	IsActive     bool   `gorm:"not null"`
	IsBlocked    bool   `gorm:"not null"`
	IsVerified   bool   `gorm:"not null"`
	IsDeleted    bool   `gorm:"not null"`
	MaxFrequency int
}

func CreateUserTable(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
}

func CreateUser(db *gorm.DB, user User) {
	db.Create(&user)
}

func GetUserByEmail(db *gorm.DB, email string) (User, error) {
	var user User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}

func GetUserByName(db *gorm.DB, name string) (User, error) {
	var user User
	err := db.Where("name = ?", name).First(&user).Error
	return user, err
}

func GetUserByID(db *gorm.DB, id int) (User, error) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	return user, err
}
