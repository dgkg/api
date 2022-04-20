package moke

import (
	"errors"

	"github.com/dgkg/api/model"
)

func (db *DB) CreateUser(u *model.User) (*model.User, error) {
	id := db.generateID(5)
	u.ID = id
	db.userList[id] = u
	return u, nil
}

func (db *DB) DeleteUser(uuiUser string) error {
	_, ok := db.userList[uuiUser]
	if !ok {
		return errors.New("db: user don't exists")
	}
	delete(db.userList, uuiUser)
	return nil
}

func (db *DB) GetAllUsers() (map[string]*model.User, error) {
	return db.userList, nil
}

func (db *DB) GetUserByID(uuidUser string) (*model.User, error) {
	u, ok := db.userList[uuidUser]
	if !ok {
		return nil, errors.New("db: user don't exists")
	}
	return u, nil
}

func (db *DB) UpdateUser(uuidUser string, data map[string]interface{}) (*model.User, error) {
	u, ok := db.userList[uuidUser]
	if !ok {
		return nil, errors.New("db: user don't exists")
	}

	updateString(&u.FirstName, "first_name", data)
	updateString(&u.LastName, "last_name", data)
	updateString(&u.Email, "email", data)

	return u, nil
}

func updateString(field *string, key string, data map[string]interface{}) {
	v, ok := data[key]
	if ok {
		value, ok := v.(string)
		if ok {
			*field = value
		}
	}
}
