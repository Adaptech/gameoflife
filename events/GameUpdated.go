package events

type GameUpdated struct {
	GameId string `json:"gameId"`
	Grid string `json:"grid"`
	
}
