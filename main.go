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

	r.GET("/posts/:id", func(c *gin.Context) {
		type Params struct {
			ID uint `uri:"id" binding:"required"`
		}

		var p Params
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"id":      p.ID,
		})
	})

	r.GET("/posts/:id/comments", func(c *gin.Context) {
		type Params struct {
			ID uint `uri:"id" binding:"required"`
		}

		var p Params
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"id":      p.ID,
		})
	})

	r.GET("/posts/:id/comments/:commentId", func(c *gin.Context) {
		type Params struct {
			ID        uint `uri:"id" binding:"required"`
			CommentID uint `uri:"commentId" binding:"required"`
		}

		var p Params
		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":   "pong",
			"id":        p.ID,
			"commentId": p.CommentID,
		})
	})

	r.Run(":" + config.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
