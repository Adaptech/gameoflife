package game

import "net/http"
import "github.com/Adaptech/gameoflife/infra"
import "github.com/Adaptech/gameoflife/commands"
import domain "github.com/Adaptech/gameoflife/domain/game"


func play_handler(w http.ResponseWriter, r *http.Request, commandHandler infra.CommandHandler) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		return
	}
	command := &commands.Play {}
	if err := infra.ReadJson(r, &command); err != nil {
		w.WriteHeader(500)
		infra.WriteJson(w, map[string]interface{}{"message": err.Error()})
		return
	}
	aggregate := &domain.Game {}
	if err := commandHandler(command.GameId, aggregate, command); err != nil {
		//TODO handle validation errors
		w.WriteHeader(500)
		infra.WriteJson(w, map[string]interface{}{"message": err.Error()})
		return
	}
	infra.WriteJson(w, command)
}


func RegisterRoutes(commandHandler infra.CommandHandler) {
	http.HandleFunc("/api/v1/game/play", func (w http.ResponseWriter, r *http.Request) {
		play_handler(w, r, commandHandler)
	})
	
}
