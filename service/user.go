package service

import "go_gin_app/models"

func AddUser(user *models.User) {
	models.AddUser(*user)
}
