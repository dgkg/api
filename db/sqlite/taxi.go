package sqlite

import (
	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

func (db *DB) CreateTaxi(u *model.Taxi) (*model.Taxi, error) {
	id := uuid.New().String()
	u.ID = id
	err := db.conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *DB) DeleteTaxi(uuidTaxi string) error {
	return db.conn.Model(model.Taxi{}).Where("id = ?", uuidTaxi).Delete(model.Taxi{}).Error
}

func (db *DB) GetAllTaxis() (map[string]*model.Taxi, error) {
	var us []model.Taxi
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	mus := make(map[string]*model.Taxi, len(us))
	for k := range us {
		mus[us[k].ID] = &us[k]
	}
	return mus, nil
}

func (db *DB) GetTaxiByID(uuidTaxi string) (*model.Taxi, error) {
	var u model.Taxi
	err := db.conn.Where("id = ?", uuidTaxi).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *DB) UpdateTaxi(uuidTaxi string, data map[string]interface{}) (*model.Taxi, error) {
	err := db.conn.Model(model.Taxi{}).Where("id = ?", uuidTaxi).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetTaxiByID(uuidTaxi)
}

var ts = map[string]*model.Taxi{
	"b9b8ec20-bff0-11ec-bc0f-33135681d549": {
		ID: "b9b8ec20-bff0-11ec-bc0f-33135681d549",
	},
	"da7bd99a-bff0-11ec-87a9-b7bdf1760731": {
		ID: "da7bd99a-bff0-11ec-87a9-b7bdf1760731",
	},
}
