package model

import "time"

type Taxi struct {
	ID          string    `json:"id" gorm:"primaryKey"`
	License     string    `json:"lisense"`
	NumberPlate string    `json:"number_plate"`
	CreateAt    time.Time `json:"create_at"`
	UpdateAt    time.Time `json:"update_at"`
}
