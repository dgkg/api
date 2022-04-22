package sqlite

import (
	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

func (db *DB) CreateLocation(u *model.Location) (*model.Location, error) {
	id := uuid.New().String()
	u.ID = id
	err := db.conn.Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (db *DB) DeleteLocation(uuidLocation string) error {
	return db.conn.Model(model.Location{}).Where("id = ?", uuidLocation).Delete(model.Location{}).Error
}

func (db *DB) GetAllLocations() (map[string]*model.Location, error) {
	var us []model.Location
	err := db.conn.Find(&us).Error
	if err != nil {
		return nil, err
	}
	mus := make(map[string]*model.Location, len(us))
	for k := range us {
		mus[us[k].ID] = &us[k]
	}
	return mus, nil
}

func (db *DB) GetLocationByID(uuidLocation string) (*model.Location, error) {
	var u model.Location
	err := db.conn.Where("id = ?", uuidLocation).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *DB) UpdateLocation(uuidLocation string, data map[string]interface{}) (*model.Location, error) {
	err := db.conn.Model(model.Location{}).Where("id = ?", uuidLocation).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return db.GetLocationByID(uuidLocation)
}
