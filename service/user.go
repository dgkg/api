package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/dgkg/api/model"
	"github.com/dgkg/api/session"
)

func (s *Service) GetAllUsers(ctx *gin.Context) {
	usdb, err := s.db.GetAllUsers()
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	users := make([]*model.User, len(usdb))
	var i int
	for k := range usdb {
		users[i] = usdb[k]
		i++
	}
	ctx.JSON(200, users)
}

func (s *Service) GetUserByID(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, nil)
		return
	}
	ctx.JSON(200, u)
}

func (s *Service) CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	_, err = s.db.CreateUser(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, u)
}

func (s *Service) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	_, err := uuid.Parse(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = s.db.DeleteUser(id)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, nil)
}

func (s *Service) UpdateUser(ctx *gin.Context) {
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

	u, err := s.db.UpdateUser(id, payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(200, u)
}

func (s *Service) LoginUser(ctx *gin.Context) {

	var payload model.Login
	err := ctx.BindJSON(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	u, err := s.db.GetUserByEmail(payload.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if u.Password != payload.Password {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := session.New(u.ID, u.AccessLevel)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(200, gin.H{
		"jwt": jwtValue,
	})
}
