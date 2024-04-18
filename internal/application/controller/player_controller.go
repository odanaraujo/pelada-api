package controller

import (
	"encoding/json"
	"fmt"
	"github.com/odanaraujo/pelada-api/internal/application/dto"
	"github.com/odanaraujo/pelada-api/internal/application/usecase"
	"net/http"
)

type Playercontroller struct {
	AddPlayerUsecase *usecase.AddPlayerUsecase
}

func NewPlayerController(addPlayerUsecase *usecase.AddPlayerUsecase) *Playercontroller {
	return &Playercontroller{AddPlayerUsecase: addPlayerUsecase}
}

func (pc *Playercontroller) AddPlayerController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("init controller")
	var requestData dto.AddPlayersListDTO
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	if err := pc.AddPlayerUsecase.Execute(r.Context(), requestData); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
}
