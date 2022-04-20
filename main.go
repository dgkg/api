package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dgkg/api/db/sqlite"
	"github.com/dgkg/api/service"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("view/*")
	r.Static("/static", "./static")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	conn := sqlite.New("mystorage.db")
	s := service.New(conn)
	r.GET("/users", s.GetAllUsers)
	r.GET("/users/:id", s.GetUserByID)
	r.POST("/users", s.CreateUser)
	r.DELETE("/users/:id", s.DeleteUser)
	r.PATCH("/users/:id", s.UpdateUser)
	r.Run() // listen and serve on 0.0.0.0:8080
}
