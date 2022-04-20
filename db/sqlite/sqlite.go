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

	return &DB{
		conn: db,
	}
}
