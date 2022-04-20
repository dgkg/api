package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dgkg/api/db"
	"github.com/dgkg/api/model"
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

func (s *Service) GetUserByID(c *gin.Context) {
	id := c.Param("id")
	if len(id) > 200 {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	_, ok := db.US[id]
	if !ok {
		c.JSON(http.StatusNotFound, nil)
		return
	}
	c.JSON(200, db.US[id])
}

func (s *Service) CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.BindJSON(&u)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
}
