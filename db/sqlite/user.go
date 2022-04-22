package sqlite

import (
	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

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

func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := db.conn.Where("email = ?", email).First(&u).Error
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
