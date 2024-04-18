package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/odanaraujo/pelada-api/internal/application/dto"
	"github.com/odanaraujo/pelada-api/internal/domain/entity"
)

type AddPlayerUsecase struct{}

func NewAddPlayerUsecase() *AddPlayerUsecase {
	return &AddPlayerUsecase{}
}

func (apu *AddPlayerUsecase) Execute(ctx context.Context, input dto.AddPlayersListDTO) error {

	playersMonthlyList := make([]entity.PlayersMonthly, 0)
	playersDiaryList := make([]entity.PlayersDiary, 0)

	for _, value := range input.Players {
		switch value.PlayerType {
		case entity.DiaryPlayersList:
			playerDiarist := entity.PlayersDiary{Name: value.Name, Age: value.Age}
			playersDiaryList = append(playersDiaryList, playerDiarist)
		case entity.MonthlyPlayersList:
			playerMonthly := entity.PlayersMonthly{Name: value.Name, Age: value.Age}
			playersMonthlyList = append(playersMonthlyList, playerMonthly)
		default:
			return errors.New("not a valid player type")
		}
	}

	playerEntity, err := entity.NewPlayers(playersMonthlyList, playersDiaryList)

	if err != nil {
		return err
	}

	fmt.Println(playerEntity)

	return nil
}
