package moke

import (
	"errors"

	"github.com/dgkg/api/model"
)

func (db *DB) CreateLocation(u *model.Location) (*model.Location, error) {
	id := db.generateID(5)
	u.ID = id
	db.locationList[id] = u
	return u, nil
}

func (db *DB) DeleteLocation(uuiLocation string) error {
	_, ok := db.locationList[uuiLocation]
	if !ok {
		return errors.New("db: Location don't exists")
	}
	delete(db.locationList, uuiLocation)
	return nil
}

func (db *DB) GetAllLocations() (map[string]*model.Location, error) {
	return db.locationList, nil
}

func (db *DB) GetLocationByID(uuidLocation string) (*model.Location, error) {
	u, ok := db.locationList[uuidLocation]
	if !ok {
		return nil, errors.New("db: Location don't exists")
	}
	return u, nil
}

func (db *DB) UpdateLocation(uuidLocation string, data map[string]interface{}) (*model.Location, error) {
	u, ok := db.locationList[uuidLocation]
	if !ok {
		return nil, errors.New("db: Location don't exists")
	}

	return u, nil
}
