package controllers

import "github.com/Adaptech/gameoflife/infra"
import "github.com/Adaptech/gameoflife/controllers/game"

func RegisterRoutes(commandHandler infra.CommandHandler) {
	game.RegisterRoutes(commandHandler)
	
}
