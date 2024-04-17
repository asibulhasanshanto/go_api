package services

import (
	"github.com/asibulhasanshanto/restapi/model"
)

func CreateUser(user *model.User) *model.User {
	result := dbConnection.Db.Create(user)

	if result.Error != nil {
		panic(result.Error)
	}

	return user
}
