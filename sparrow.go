package main

import (
	"github.com/soulteary/sparrow/internal/server"
)

// var brokerPool = br.NewManager(256)

func main() {
	engine := server.SetupEngine()
	// api.AuthSession(engine)
	// api.Public(engine)
	// api.Backend(engine, brokerPool)
	server.Launched(engine)
}
