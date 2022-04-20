package service

import (
	"github.com/dgkg/api/db/sqlite"
)

type Service struct {
	db *sqlite.DB
}

func New(db *sqlite.DB) *Service {
	return &Service{
		db: db,
	}
}
