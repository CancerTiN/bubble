package router

import (
	"bubble/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() (engine *gin.Engine) {
	engine = gin.Default()

	engine.Static("/static", "static")
	engine.LoadHTMLGlob("templates/*")
	engine.GET("/", controller.IndexHandler)

	v1Group := engine.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateTodo)
		v1Group.GET("/todo", controller.GetTodoSlice)
		v1Group.PUT("/todo/:id", controller.UpdateATodo)
		v1Group.DELETE("/todo/:id", controller.DeleteATodo)
	}
	return
}
