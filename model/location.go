package model

import (
	"time"
)

type Location struct {
	ID       string    `json:"id"`
	IDOwner  string    `json:"id_owner"`
	Lat      float64   `json:"lat"`
	Long     float64   `json:"long"`
	CreateAt time.Time `json:"create_at"`
}
