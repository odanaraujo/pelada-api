package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odanaraujo/pelada-api/cmd/api/handler"
	"github.com/odanaraujo/pelada-api/internal/application/controller"
	"github.com/odanaraujo/pelada-api/internal/application/usecase"
	"log"
)

func main() {
	addPlayerUsecase := usecase.NewConfirmePresenceUsecase()
	playercontroller := controller.NewPlayerController(addPlayerUsecase)

	r := gin.Default()
	handler.InitHandler(&r.RouterGroup, *playercontroller)

	if err := r.Run(); err != nil {
		log.Printf("error initialize server %v", err.Error())
		return
	}
}
