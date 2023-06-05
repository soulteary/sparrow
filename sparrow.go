package main

import (
	"fmt"

	eb "github.com/soulteary/sparrow/components/event-broker"
	lcs "github.com/soulteary/sparrow/components/local-conversation-storage"
	claude "github.com/soulteary/sparrow/connectors/claude"
	midjourney "github.com/soulteary/sparrow/connectors/mid-journey"
	"github.com/soulteary/sparrow/internal/api"
	"github.com/soulteary/sparrow/internal/define"
	"github.com/soulteary/sparrow/internal/server"
	"github.com/soulteary/sparrow/internal/version"
)

var brokerPool = eb.NewBrokerManager(256)

func main() {
	fmt.Printf("Sparrow v%s\n", version.Version)

	lcs.InitStorage(define.LOCAL_CONVERSATION_STORAGE_PATH)

	engine := server.SetupEngine()
	engine.Use(server.Recovery())

	api.AuthSession(engine)
	api.Public(engine)
	api.Backend(engine, brokerPool)

	if define.ENABLE_MIDJOURNEY {
		go midjourney.KeepConnection(brokerPool)
	}

	if define.ENABLE_CLAUDE {
		go claude.KeepConnection(brokerPool)
	}

	server.Launched(engine)
}
