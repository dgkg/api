package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dgkg/api/db"
	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

var _ db.Storage = &DB{}

type DB struct {
	conn *gorm.DB
}

func New(dbName string) *DB {
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.User{})
	for _, u := range US {
		db.Create(u)
	}

	return &DB{
		conn: db,
	}
}

func (db *DB) CreateUser(u *model.User) (*model.User, error) {
	id := uuid.New().String()
	u.ID = id
	err := db.conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *DB) DeleteUser(uuidUser string) error {
	return db.conn.Model(model.User{}).Where("id = ?", uuidUser).Delete(model.User{}).Error
}

func (db *DB) GetAllUsers() (map[string]*model.User, error) {
	var us []model.User
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	mus := make(map[string]*model.User, len(us))
	for k := range us {
		mus[us[k].ID] = &us[k]
	}
	return mus, nil
}

func (db *DB) GetUserByID(uuidUser string) (*model.User, error) {
	var u model.User
	err := db.conn.Where("id = ?", uuidUser).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *DB) UpdateUser(uuidUser string, data map[string]interface{}) (*model.User, error) {
	err := db.conn.Model(model.User{}).Where("id = ?", uuidUser).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetUserByID(uuidUser)
}

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
