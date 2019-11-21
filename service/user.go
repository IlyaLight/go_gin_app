package service

import (
	"go_gin_app/models"
	"go_gin_app/util"
)

func AddUser(user *models.User) {
	//>проверка
	util.PasswordHashing(user)
	models.AddUser(user)
}

func VerifyUser(username, password string) (*models.User, error) {
	//user, err := models.GetUserByUsername(username)
	//if err != nil{
	//	return nil, err
	//}
	//if user == nil{
	//	return nil, errors.New(fmt.Sprintf("User with username: %s doesn't exists", username))
	//}
	//if result, _ := util.ComparePasswords(user, password); result != true{
	//	return nil, errors.New("password doesnt't match")
	//}
	//util.GenerateToken(user)
	////util.CheckToken(&user.Token)
	//return user, err
	return nil, nil
}
