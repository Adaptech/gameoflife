package readModels

import "github.com/Adaptech/gameoflife/infra"
import "github.com/Adaptech/gameoflife/readModels/grid"

func GetHandler(eventType string) infra.EventHandler {
	switch eventType {
	case "grid":
		return grid.Handler
	
	}
	return nil
}
