package fuzz

import "github.com/dgkg/api/model"

type Data struct {
	LocationList map[string]*model.Location
	UserList     map[string]*model.User
	TaxiList     map[string]*model.Taxi
}

func New(nbTaxi, nbUser int) (*Data, error) {
	var res Data
	var err error
	res.TaxiList, err = FuzzTaxi(nbTaxi)
	if err != nil {
		return nil, err
	}
	res.UserList, err = FuzzUser(nbUser)
	if err != nil {
		return nil, err
	}

	res.LocationList, err = FuzzLocation(res.UserList, res.TaxiList)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
