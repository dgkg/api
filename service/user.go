package service

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dgkg/api/db"
	"github.com/dgkg/api/model"
)

func GetUsers(c *gin.Context) {
	users := make([]model.User, len(db.US))
	var i int
	for k := range db.US {
		users[i] = db.US[k]
		i++
	}
	c.JSON(200, users)
}

func GetUserByID(c *gin.Context) {
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
