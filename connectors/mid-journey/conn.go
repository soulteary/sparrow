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

func GetConn() *websocket.Conn {
	if Conn != nil {
		return Conn
	} else {
		log.Println("Connected to Midjoruney API server.")
		Conn = GetWebSocketConn(define.MODJOURNEY_API_SECRET)
		return Conn
	}
}

func GetWebSocketConn(secret string) *websocket.Conn {
	u, _ := url.Parse(fmt.Sprintf("%s/%s", define.MIDJOURNEY_API_ADDR, secret))
	log.Printf("connecting to %s", u.String())

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	return conn
}

func KeepConnection(brokerPool *eb.BrokersPool) {
	conn := GetConn()
	defer conn.Close()
	done := make(chan bool)
	CreateReceiver(&done, conn, brokerPool, FnReceiver())
}
