package grid

import "github.com/Adaptech/gameoflife/infra"
import "github.com/Adaptech/gameoflife/events"

type Grid struct {
	GameId string `json:"gameId"`
	Grid string `json:"grid"`
	
}

//TODO PK gameId

func Handler(gridRepo *infra.TrxRepository, eventData *infra.EventData) error {
	event := eventData.Event
	record := Grid {}
	
	if ev, ok := event.(*events.GameUpdated); ok {
		record.GameId = ev.GameId
		record.Grid = ev.Grid
		
	}
	
	return gridRepo.Create(record)
}
