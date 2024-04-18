package main

import (
	"errors"
	"fmt"
	"github.com/odanaraujo/pelada-api/internal/application/controller"
	"github.com/odanaraujo/pelada-api/internal/application/usecase"
	"net/http"
)

func main() {
	addPlayerUsecase := usecase.AddPlayerUsecase{}
	playercontroller := controller.NewPlayerController(&addPlayerUsecase)

	http.HandleFunc("/add-players", playercontroller.AddPlayerController)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		errors.New(fmt.Sprintf("Unable to initialize server %v", err.Error()))
	}

}
