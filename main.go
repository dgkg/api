package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
	r.GET("/users", service.GetUsers)
	r.GET("/users/:id", service.GetUserByID)
	r.Run() // listen and serve on 0.0.0.0:8080
}
