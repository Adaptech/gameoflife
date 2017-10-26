package infra

import (
	"github.com/jdextraze/go-gesclient/client"
)

type Infra struct {
	eventFactory             EventFactory
	readModelHandlerProvider EventHandlerProvider
	registerRoutes           func(CommandHandler)
	userCredentials		 	 *client.UserCredentials
	conn                     client.Connection
}

func New() *Infra {
	return &Infra{}
}

func (b *Infra) UsingEventFactory(eventFactory EventFactory) *Infra {
	b.eventFactory = eventFactory
	return b
}

func (b *Infra) UsingReadModelHandlerProvider(readModelHandlerProvider EventHandlerProvider) *Infra {
	b.readModelHandlerProvider = readModelHandlerProvider
	return b
}

func (b *Infra) UsingRegisterRoutes(registerRoutes func(CommandHandler)) *Infra {
	b.registerRoutes = registerRoutes
	return b
}

func (b *Infra) UsingUserCredentials(username string, password string) *Infra {
	b.userCredentials = client.NewUserCredentials(username, password)
	return b
}
