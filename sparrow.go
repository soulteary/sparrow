package main

import (
	"fmt"

	eb "github.com/soulteary/sparrow/components/event-broker"
	"github.com/soulteary/sparrow/internal/api"
	"github.com/soulteary/sparrow/internal/server"
	"github.com/soulteary/sparrow/internal/version"
)

var brokerPool = eb.NewBrokerManager(256)

func main() {
	fmt.Printf("Sparrow v%s\n", version.Version)

	engine := server.SetupEngine()
	engine.Use(server.Recovery())
	api.AuthSession(engine)
	api.Public(engine)
	api.Backend(engine, brokerPool)
	server.Launched(engine)
}
