package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dgkg/api/model"
)

func (s *Service) GetAllLocations(ctx *gin.Context) {
	usdb, err := s.db.GetAllLocations()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	Location := make([]*model.Location, len(usdb))
	var i int
	for k := range usdb {
		Location[i] = usdb[k]
		i++
	}
	ctx.JSON(200, Location)
}

func (s *Service) GetLocationByID(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.GetLocationByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(200, u)
}

func (s *Service) CreateLocation(ctx *gin.Context) {
	var u model.Location
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = s.db.CreateLocation(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, u)
}

func (s *Service) DeleteLocation(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = s.db.DeleteLocation(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, nil)
}

func (s *Service) UpdateLocation(ctx *gin.Context) {
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

	u, err := s.db.UpdateLocation(id, payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, u)
}
