package infra

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/jdextraze/go-gesclient"
	"github.com/jdextraze/go-gesclient/client"
)

func (i *Infra) Run() error {
	if i.eventFactory == nil {
		log.Fatalf("Missing eventFactory")
	}
	if i.readModelHandlerProvider == nil {
		log.Fatalf("Missing readModelHandlerProvider")
	}
	if i.registerRoutes == nil {
		log.Fatalf("Missing registerRoutes")
	}
	if i.userCredentials == nil {
		i.userCredentials = client.NewUserCredentials("admin", "changeit")
	}

	gesAddr := "tcp://eventstore:1113"
	if uri, err := url.Parse(gesAddr); err != nil {
		log.Fatalf("Wrong format for GES Address: %s", err.Error())
	} else if i.conn, err = gesclient.Create(nil, uri, "conn"); err != nil {
		log.Fatalf("Error creating connection for GES: %s", err.Error())
	}

	i.conn.Closed().Add(func(event client.Event) error {
		log.Println("Connection to GES lost, shutting down app.")
		os.Exit(0)
		return nil
	})

	log.Println("Connecting to GES at ", gesAddr)
	if err := i.conn.ConnectAsync().Wait(); err != nil {
		log.Fatalf("Error connecting to GES: %s", err.Error())
	}

	http.Handle("/api/v1/r/", i.NewGenericReadModelsHandler())

	i.registerRoutes(i.HandleCommand)

	if err := http.ListenAndServe(":3001", nil); err != nil {
		log.Fatalf("Error starting web server: %s", err.Error())
	}

	return nil
}
