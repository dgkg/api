package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/dgkg/api/db"
	"github.com/dgkg/api/db/moke"
	"github.com/dgkg/api/db/sqlite"
	"github.com/dgkg/api/service"
)

type Config struct {
	Port   string
	DBName string
	Env    string
}

var config Config

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", err))
	}
	log.Println("ENV type:", viper.GetString("ENV"))
	log.Println("DBName:", viper.GetString("DBName"))
	config.DBName = viper.GetString("DBName")
	config.DBName = viper.GetString("DBName")
	config.DBName = viper.GetString("Port")
}

const (
	EnvProduction = "production"
	EnvLocal      = "local"
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
	var conn db.Storage
	if viper.GetString("ENV") == EnvLocal {
		conn = moke.New()
	} else {
		conn = sqlite.New("mystorage.db")
	}

	service.New(r, conn)
	r.Run(":" + config.Port)
}
