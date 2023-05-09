package main

import (
	eb "github.com/soulteary/sparrow/components/event-broker"
	"github.com/soulteary/sparrow/internal/api"
	"github.com/soulteary/sparrow/internal/server"
)

var brokerPool = eb.NewBrokerManager(256)

func main() {
	engine := server.SetupEngine()
	api.AuthSession(engine)
	api.Public(engine)
	api.Backend(engine, brokerPool)
	server.Launched(engine)
}
