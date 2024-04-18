package main

import (
	"errors"
	"fmt"
	"github.com/odanaraujo/pelada-api/internal/application/controller"
	"github.com/odanaraujo/pelada-api/internal/application/usecase"
	"net/http"
)

func main() {
	addPlayerUsecase := usecase.ConfirmePresenceUsecase{}
	playercontroller := controller.NewPlayerController(&addPlayerUsecase)

	http.HandleFunc("/confirm-presence", playercontroller.ConfirmePresenceController)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		errors.New(fmt.Sprintf("Unable to initialize server %v", err.Error()))
	}

}
