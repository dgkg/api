package db

import "github.com/dgkg/api/model"

type DB struct {
	userList map[string]*model.User
}

func New() *DB {
	return &DB{
		userList: US,
	}
}

// func CreateUser
// func DeleteUser
// func GetUserByID
// func UpdateUser

var US = map[string]*model.User{
	"b9b8ec20-bff0-11ec-bc0f-33135681d549": {
		ID:        "b9b8ec20-bff0-11ec-bc0f-33135681d549",
		FirstName: "Bob",
		LastName:  "Willson",
		Email:     "b@willson.com",
	},
	"da7bd99a-bff0-11ec-87a9-b7bdf1760731": {
		ID:        "da7bd99a-bff0-11ec-87a9-b7bdf1760731",
		FirstName: "Andrea",
		LastName:  "Guisberg",
		Email:     "a@guisberg.co.uk",
	},
}
