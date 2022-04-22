package fuzz

import (
	"errors"
	"math/rand"

	"github.com/dgkg/api/model"
	"github.com/google/uuid"
)

func FuzzTaxi(n int) (map[string]*model.Taxi, error) {
	if n > 255000 {
		return nil, errors.New("fuzz: out of boundaries max 255 000 users")
	}

	ts := make(map[string]*model.Taxi, n)
	for i := 0; i < n; i++ {
		id := uuid.New().String()
		ts[id] = &model.Taxi{
			ID:          id,
			NumberPlate: getChartPlateNumber(),
		}
	}
	return ts, nil
}

var charPlateNumber = "QWERTYUIOPLKJHGFDSAMNBVCXZ1234567890"

func getChartPlateNumber() string {
	res := ""
	for i := 0; i < 6; i++ {
		res += string(charPlateNumber[rand.Intn(len(charPlateNumber))])
	}
	return res
}
