package models

import "github.com/jinzhu/gorm"

type User struct {
	ID       int       `json:"id"			gorm:"primary_key"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Article  []Article `gorm:"ForeignKey:UserId"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var user User
	err := db.Select("id").Where(User{Username: username, Password: password}).First(&user).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if user.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddUser(user User) {
	db.Create(&user)
}
