package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dgkg/api/db"
	"github.com/dgkg/api/model"
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
	for _, u := range us {
		db.Create(u)
	}

	db.AutoMigrate(&model.Taxi{})
	for _, t := range ts {
		db.Create(t)
	}

	db.AutoMigrate(&model.Location{})
	for _, l := range ls {
		db.Create(l)
	}

	return &DB{
		conn: db,
	}
}

var us = map[string]*model.User{
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

var ls = map[string]*model.Location{
	"67284df4-c0bb-11ec-b8ce-cb4b09959b2": {
		ID: "67284df4-c0bb-11ec-b8ce-cb4b09959b2",
	},
	"70018c56-c0bb-11ec-86a1-93ca448bf895": {
		ID: "70018c56-c0bb-11ec-86a1-93ca448bf895",
	},
}
