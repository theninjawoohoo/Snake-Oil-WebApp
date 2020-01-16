package main

import (
	"log"
	"os"
	"./handler/getIndex.go"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*")
	router.Static("/static", "./static")

	router.GET("/", handler.getIndex())
	router.POST("/game", handler.gameGet())
	//router.POST("/game", handler.getPost())

	router.Run(":" + port)
}
