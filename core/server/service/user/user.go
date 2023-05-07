package user

import (
	"strings"
	"wordie/core/db"
	"wordie/core/db/userDB"
)

type Info struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetUserFrequency(email string) (int, error) {
	database := db.Instance()
	var user userDB.User
	err := database.Where("email = ?", email).First(&user).Error
	if err != nil {
		return 0, err
	}
	return user.MaxFrequency, nil
}

func Register(_user Info) error {
	user := &userDB.User{
		Name:       _user.Username,
		Email:      _user.Email,
		Password:   _user.Password,
		IsAdmin:    false,
		IsActive:   true,
		IsBlocked:  false,
		IsVerified: false,
		IsDeleted:  false,
	}
	//regulate user email to lowercase
	user.Email = strings.ToLower(user.Email)
	database := db.Instance()
	err := database.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func VerifyUser(email string, password string) (userDB.User, error) {
	database := db.Instance()
	userToVerify := userDB.User{}
	res := database.Where("email = ?", strings.ToLower(email)).First(&userToVerify)
	if res.Error != nil {
		if userToVerify.Password == password {
			return userToVerify, nil
		}
		return userDB.User{}, res.Error
	}
	return userToVerify, nil
}
