package service

import (
	"github.com/dgkg/api/db"
	"github.com/gin-gonic/gin"
)

type Service struct {
	db db.Storage
}

func New(r *gin.Engine, db db.Storage) *Service {
	s := &Service{
		db: db,
	}
	// users
	r.GET("/users", s.GetAllUsers)
	r.GET("/users/:id", s.GetUserByID)
	r.POST("/users", s.CreateUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.PATCH("/users/:id", s.UpdateUser)

	// taxis
	r.GET("/users", s.GetAllUsers)
	r.GET("/users/:id", s.GetUserByID)
	r.POST("/users", s.CreateUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.PATCH("/users/:id", s.UpdateUser)
	return s
}
