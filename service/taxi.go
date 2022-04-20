package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dgkg/api/model"
)

func (s *Service) GetAllTaxi(ctx *gin.Context) {
	usdb, err := s.db.GetAllTaxis()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	taxi := make([]*model.Taxi, len(usdb))
	var i int
	for k := range usdb {
		taxi[i] = usdb[k]
		i++
	}
	ctx.JSON(200, taxi)
}

func (s *Service) GetTaxiByID(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.GetTaxiByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(200, u)
}

func (s *Service) CreateTaxi(ctx *gin.Context) {
	var u model.Taxi
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = s.db.CreateTaxi(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, u)
}

func (s *Service) DeleteTaxi(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = s.db.DeleteTaxi(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, nil)
}

func (s *Service) UpdateTaxi(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	payload := make(map[string]interface{})
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.UpdateTaxi(id, payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, u)
}
