package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	Config "github.com/tayalone/ma-go-wire/config"
)

func main() {
	fmt.Println("Start ma go wire")
	config := Config.Initialize()
	fmt.Println("config.Port", config.Port)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.Run(":" + config.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
