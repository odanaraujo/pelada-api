package entity

const (
	MonthlyPlayersList = "MonthlyPlayers"
	DiaryPlayersList   = "DiaryPlayers"
)

type Players struct {
	playersMonthly []PlayersMonthly
	diaryPlayers   []PlayersDiary
}

type PlayersMonthly struct {
	Name string
	Age  int
}

type PlayersDiary struct {
	Name string
	Age  int
}

func NewPlayers(playersMonthly []PlayersMonthly, playersDiary []PlayersDiary) (*Players, error) {
	return &Players{
		playersMonthly: playersMonthly,
		diaryPlayers:   playersDiary,
	}, nil
}

func (p *Players) GetPlayersMonthly() []PlayersMonthly {
	return p.playersMonthly
}

func (p *Players) GetPlayersDiary() []PlayersDiary {
	return p.diaryPlayers
}
