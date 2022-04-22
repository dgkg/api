package db

import (
	"github.com/dgkg/api/model"
)

type Storage interface {
	StorageUser
	StorageTaxi
	StorageLocation
}

type StorageUser interface {
	CreateUser(u *model.User) (*model.User, error)
	DeleteUser(uuiUser string) error
	GetAllUsers() (map[string]*model.User, error)
	GetUserByID(uuidUser string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	UpdateUser(uuidUser string, data map[string]interface{}) (*model.User, error)
}

type StorageTaxi interface {
	CreateTaxi(u *model.Taxi) (*model.Taxi, error)
	DeleteTaxi(uuiTaxi string) error
	GetAllTaxis() (map[string]*model.Taxi, error)
	GetTaxiByID(uuidTaxi string) (*model.Taxi, error)
	UpdateTaxi(uuidTaxi string, data map[string]interface{}) (*model.Taxi, error)
}

type StorageLocation interface {
	CreateLocation(u *model.Location) (*model.Location, error)
	DeleteLocation(uuiLocation string) error
	GetAllLocations() (map[string]*model.Location, error)
	GetLocationByID(uuidLocation string) (*model.Location, error)
	UpdateLocation(uuidLocation string, data map[string]interface{}) (*model.Location, error)
}
