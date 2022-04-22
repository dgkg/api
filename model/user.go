package model

type User struct {
	ID          string `json:"id" gorm:"primaryKey"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessLevel int    `json:"access_level"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
