package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type User struct {
	ID          string   `json:"id" gorm:"primaryKey"`
	FirstName   string   `json:"first_name"`
	LastName    string   `json:"last_name"`
	Email       string   `json:"email"`
	Password    Password `json:"password,omitempty"`
	AccessLevel int      `json:"access_level"`
}

type Password string

func (p *Password) UnmarshalJSON(b []byte) error {
	var ps string
	if err := json.Unmarshal(b, &ps); err != nil {
		return err
	}

	h := sha256.New()
	h.Write([]byte(ps))
	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))

	return nil
}

func (Password) MarshalJSON() ([]byte, error) {
	var s string = ""
	return json.Marshal(s)
}

type Login struct {
	Email    string   `json:"email"`
	Password Password `json:"password"`
}
