package service

import (
	"github.com/dgkg/api/db"
	"github.com/dgkg/api/session"
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
	r.GET("/users", session.NewMiddleware(1), s.GetAllUsers)
	r.GET("/users/:id", s.GetUserByID)
	r.POST("/users", s.CreateUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.PATCH("/users/:id", s.UpdateUser)
	r.POST("/login", s.LoginUser)

	// taxis
	r.GET("/taxis", s.GetAllTaxis)
	r.GET("/taxis/:id", s.GetTaxiByID)
	r.POST("/taxis", s.CreateTaxi)
	r.DELETE("/taxis/:id", s.DeleteTaxi)
	r.PATCH("/taxis/:id", s.UpdateTaxi)

	// location
	r.GET("/locations", s.GetAllLocations)
	r.GET("/locations/:id", s.GetLocationByID)
	r.POST("/locations", s.CreateLocation)
	r.DELETE("/locations/:id", s.DeleteLocation)
	r.PATCH("/locations/:id", s.UpdateLocation)
	return s
}
