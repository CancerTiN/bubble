package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status bool `json:"status"`
}

func main() {
	engine := gin.Default()

	engine.Static("/static", "static")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := engine.Group("v1")
	{
		v1Group.POST("/todo", func(c *gin.Context) {

		})
		v1Group.GET("/todo", func(c *gin.Context) {

		})
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})
		v1Group.PUT("/todo/:id", func(c *gin.Context) {

		})
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
