package main

import (
	"fuego_backend/controller"
	"fuego_backend/database"
	"fuego_backend/models"

	"github.com/go-fuego/fuego"
)

func main() {
	db := database.ConnectDB()

	err := db.AutoMigrate(&models.Server{}, &models.VM{}, &models.Container{})
	if err != nil {
		panic("AutoMigrate failed: " + err.Error())
	}

	server := fuego.NewServer()

	controller.ServerResources{
		ServerService: controller.NewPersistentServerService(db),
	}.Routes(server)

	server.Run()
}
