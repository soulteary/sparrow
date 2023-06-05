package Claude

import (
	"fmt"
	"log"
	"strings"

	"github.com/gorilla/websocket"
	eb "github.com/soulteary/sparrow/components/event-broker"
	sr "github.com/soulteary/sparrow/components/stream-responser"
	"github.com/soulteary/sparrow/internal/define"
)

func PostMessage(conn *websocket.Conn, message []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		log.Println("write:", err)
		return err
	}
	return nil
}

func CreateReceiver(done *chan bool, conn *websocket.Conn, brokerPool *eb.BrokersPool, fn func(error, []byte, *eb.BrokersPool)) {
	for {
		select {
		case <-*done:
			fn(fmt.Errorf("stop messages receiver"), nil, brokerPool)
			return
		default:
			messageType, payload, err := conn.ReadMessage()
			if err != nil {
				fn(err, payload, brokerPool)
				return
			}
			if messageType != websocket.TextMessage {
				fn(fmt.Errorf("not text message"), payload, brokerPool)
				return
			}

			fn(nil, payload, brokerPool)
		}
	}
}

func ParseMessage(payload []byte) (parentMessageID string, conversationID string, response string, done bool, err error) {
	input := string(payload)
	texts := strings.Split(input, "\n")

	if define.DEV_MODE {
		fmt.Println("[Message content]")
		fmt.Println(input)
	}

	if len(texts) < 2 {
		fmt.Println(texts)
		return "", "", "", false, fmt.Errorf("the format of the message sent back by the server is incorrect")
	}

	parentMessageID = strings.TrimSpace(texts[0])
	conversationID = ""
	ControlText := strings.TrimSpace(texts[1])
	response = strings.Join(texts[2:], "\n")
	if ControlText == "[MESSAGE:CLOSE]" {
		return parentMessageID, conversationID, response, true, nil
	}
	return parentMessageID, conversationID, response, false, nil
}

func FnReceiver() func(err error, p []byte, brokerPool *eb.BrokersPool) {
	return func(err error, p []byte, brokerPool *eb.BrokersPool) {
		if err != nil {
			fmt.Println("Error receiving message", err)
			return
		}

		parentMessageID, conversationID, response, done, err := ParseMessage(p)
		if err != nil {
			fmt.Println("Error parsing message", err)
			return
		}

		modelSlug := "claude"

		user := "user"
		broker := brokerPool.GetBroker(user, parentMessageID, conversationID)
		if broker == nil {
			fmt.Println("Unable to find Broker by ParentMessageID", parentMessageID)
			return
		}

		if !done {
			sr.StreamBuilder(user, parentMessageID, conversationID, modelSlug, broker, response, sr.MSG_STATUS_CONTINUE)
		} else {
			sr.StreamBuilder(user, parentMessageID, conversationID, modelSlug, broker, response, sr.MSG_STATUS_DONE)
		}
	}
}
