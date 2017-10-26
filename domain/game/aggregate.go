package game

import "errors"
import "github.com/Adaptech/gameoflife/commands"
import "github.com/Adaptech/gameoflife/events"

type Game struct {
	GameId string
	Grid string
	
}

func (a Game) Hydrate(event interface{}) {
	if ev, ok := event.(*events.GameUpdated); ok {
		a.onGameUpdated(ev)
	}
	
}


func (a *Game) onGameUpdated(event *events.GameUpdated) {
	a.GameId = event.GameId
	a.Grid = event.Grid
	
}


func (a Game) Execute(command interface {}) ([]interface{}, error) {
	if cmd, ok := command.(*commands.Play); ok {
		return a.Play(cmd);
	}
	
	//TODO reflect command name?
	return nil, errors.New("Unknown command for Game.")
}


func (a *Game) Play(command *commands.Play) ([]interface{}, error) {
	// TODO: Validation Errors
	//const validationErrors = [];
	//if(validationErrors.length > 0) {
	//	throw new errors.ValidationFailed(validationErrors);
	//}

	var result []interface{}
	result = append(result, &events.GameUpdated {
		command.GameId, command.Grid, 
	})
	return result, nil
}

