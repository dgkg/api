package model

type Taxi struct {
	ID          string `json:"id" gorm:"primaryKey"`
	License     string `json:"lisense"`
	NumberPlate string `json:"number_plate"`
	CreateAt    int64  `json:"create_at"`
	UpdateAt    int64  `json:"update_at"`
}
