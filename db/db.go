package db

import (
	"github.com/dgkg/api/model"
)

type Storage interface {
	CreateUser(u *model.User) (*model.User, error)
	DeleteUser(uuiUser string) error
	GetAllUsers() (map[string]*model.User, error)
	GetUserByID(uuidUser string) (*model.User, error)
	UpdateUser(uuidUser string, data map[string]interface{}) (*model.User, error)
}
