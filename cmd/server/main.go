package main

import (
	"TaoCoupon/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.InitConfig()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
