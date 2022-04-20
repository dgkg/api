package service

import (
	"github.com/dgkg/api/db"
)

type Service struct {
	db db.Storage
}

func New(db db.Storage) *Service {
	return &Service{
		db: db,
	}
}
