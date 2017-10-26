package main

import "log"
import "github.com/Adaptech/gameoflife/infra"
import "github.com/Adaptech/gameoflife/events"
import "github.com/Adaptech/gameoflife/readModels"
import "github.com/Adaptech/gameoflife/controllers"

func main() {
	log.Println("Adaptech's Golang Stack")
	err := infra.New().
		UsingEventFactory(events.CreateEvent).
		UsingReadModelHandlerProvider(readModels.GetHandler).
		UsingRegisterRoutes(controllers.RegisterRoutes).
		Run()
	if err != nil {
		log.Fatalf("Error starting app: %s", err.Error())
	}
}
