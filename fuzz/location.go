package fuzz

import (
	"math/rand"

	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

func FuzzLocation(
	us map[string]*model.User,
	tx map[string]*model.Taxi) (map[string]*model.Location, error) {

	ls := make(map[string]*model.Location)

	// create taxis.
	for k := range tx {
		id := uuid.New().String()
		ls[id] = &model.Location{
			ID:      id,
			IDOwner: k,
			Lat:     float64(rand.Intn(1000)/1000 + 5),
			Long:    float64(rand.Intn(1000)/1000 + 43),
		}
	}

	// create users
	for k := range us {
		id := uuid.New().String()
		ls[id] = &model.Location{
			ID:      id,
			IDOwner: k,
			Lat:     float64(rand.Intn(1000)/1000 + 5),
			Long:    float64(rand.Intn(1000)/1000 + 43),
		}
	}

	return ls, nil
}
