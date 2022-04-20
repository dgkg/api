package moke

import (
	"github.com/dgkg/api/db"
	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

var _ db.Storage = &DB{}

type DB struct {
	userList map[string]*model.User
	taxiList map[string]*model.Taxi
}

func New() *DB {
	return &DB{
		userList: us,
		taxiList: ts,
	}
}

func (db *DB) generateID(retry int) string {
	id := uuid.New().String()
	_, ok := db.userList[id]
	if ok {
		if retry > 0 {
			id = db.generateID(retry - 1)
		} else {
			panic("not able to generate unique uuid")
		}
	}
	return id
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

var ts = map[string]*model.Taxi{
	"67284df4-c0bb-11ec-b8ce-cb4b09959b2": {
		ID: "67284df4-c0bb-11ec-b8ce-cb4b09959b2",
	},
	"70018c56-c0bb-11ec-86a1-93ca448bf895": {
		ID: "70018c56-c0bb-11ec-86a1-93ca448bf895",
	},
}
