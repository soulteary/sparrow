package midjourney

import (
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	eb "github.com/soulteary/sparrow/components/event-broker"
	"github.com/soulteary/sparrow/internal/define"
)

var Conn *websocket.Conn
var Ready bool

func GetConn(reconnect bool) *websocket.Conn {
	if Conn != nil && !reconnect {
		return Conn
	} else {
		log.Println("Connected to Midjoruney API server.")
		Conn = GetWebSocketConn(define.MIDJOURNEY_API_SECRET)
		return Conn
	}
}

func GetWebSocketConn(secret string) *websocket.Conn {
	u, _ := url.Parse(fmt.Sprintf("%s/%s", define.MIDJOURNEY_API_ADDR, secret))
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		Ready = false
		log.Fatal("dial:", err)
	}
	Ready = true
	return conn
}

func KeepConnection(brokerPool *eb.BrokersPool, reconnect bool) {
	conn := GetConn(reconnect)
	defer conn.Close()
	done := make(chan bool)
	CreateReceiver(&done, conn, brokerPool, FnReceiver())
}
