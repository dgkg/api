package moke

import (
	"errors"

	"github.com/dgkg/api/model"
)

func (db *DB) CreateTaxi(u *model.Taxi) (*model.Taxi, error) {
	id := db.generateID(5)
	u.ID = id
	db.taxiList[id] = u
	return u, nil
}

func (db *DB) DeleteTaxi(uuiTaxi string) error {
	_, ok := db.taxiList[uuiTaxi]
	if !ok {
		return errors.New("db: Taxi don't exists")
	}
	delete(db.taxiList, uuiTaxi)
	return nil
}

func (db *DB) GetAllTaxis() (map[string]*model.Taxi, error) {
	return db.taxiList, nil
}

func (db *DB) GetTaxiByID(uuidTaxi string) (*model.Taxi, error) {
	u, ok := db.taxiList[uuidTaxi]
	if !ok {
		return nil, errors.New("db: Taxi don't exists")
	}
	return u, nil
}

func (db *DB) UpdateTaxi(uuidTaxi string, data map[string]interface{}) (*model.Taxi, error) {
	u, ok := db.taxiList[uuidTaxi]
	if !ok {
		return nil, errors.New("db: Taxi don't exists")
	}

	return u, nil
}
