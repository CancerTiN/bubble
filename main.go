package main

import (
	"bubble/dao"
	"bubble/model"
	"bubble/router"
)

func main() {
	dao.DB.AutoMigrate(&model.Todo{})
	defer dao.CloseMySQL()

	engine := router.SetupRouter()

	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
