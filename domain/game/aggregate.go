package game

import "errors"
import "github.com/Adaptech/gameoflife/commands"
import "github.com/Adaptech/gameoflife/events"

type Game struct {
	GameId string
	Grid   string
}

const GRID_SIZE = 8

func (a Game) Hydrate(event interface{}) {
	if ev, ok := event.(*events.GameUpdated); ok {
		a.onGameUpdated(ev)
	}

}

func (a *Game) onGameUpdated(event *events.GameUpdated) {
	a.GameId = event.GameId
	a.Grid = event.Grid

}

func (a Game) Execute(command interface{}) ([]interface{}, error) {
	if cmd, ok := command.(*commands.Play); ok {
		return a.Play(cmd)
	}

	//TODO reflect command name?
	return nil, errors.New("unknown command for Game")
}

func (a *Game) Play(command *commands.Play) ([]interface{}, error) {
	if len(command.Grid) != 64 {
		return nil, errors.New("not a valid grid")
	}

	var result []interface{}
	result = append(result, &events.GameUpdated{
		GameId: command.GameId,
		Grid:   command.Grid,
	})
	return result, nil
}
