package EventBroker

import "fmt"

type Pool struct {
	ID              int
	ParentMessageID string
	ConversationID  string
	User            string
	Busy            bool
	Broker          Broker
}

type BrokersPool struct {
	Pool *[]Pool
}

var PoolCache []Pool

func NewBrokerManager(size int) *BrokersPool {
	for i := 0; i < size; i++ {
		broker := NewBroker(i)
		go broker.Listen()
		PoolCache = append(PoolCache, Pool{ID: i, ConversationID: "", ParentMessageID: "", User: "", Busy: false, Broker: *broker})
	}
	return &BrokersPool{Pool: &PoolCache}
}

func (broker *BrokersPool) GetBroker(user string, ParentMessageID string, ConversationID string) *Broker {
	if ParentMessageID != "" {
		for _, pool := range PoolCache {
			if pool.Busy {
				if pool.ParentMessageID == ParentMessageID {
					return &pool.Broker
				}
			}
		}
	}

	if ConversationID != "" {
		for _, pool := range PoolCache {
			if pool.Busy {
				if pool.ConversationID == ConversationID {
					return &pool.Broker
				}
			}
		}
	}

	for i := 0; i < len(PoolCache); i++ {
		if !PoolCache[i].Busy {
			PoolCache[i].Busy = true
			PoolCache[i].ConversationID = ConversationID
			PoolCache[i].ParentMessageID = ParentMessageID
			PoolCache[i].User = user
			return &PoolCache[i].Broker
		}
	}

	fmt.Println("No available broker")
	return nil
}

func FreePool(id int) {
	for i := 0; i < len(PoolCache); i++ {
		if PoolCache[i].ID == id {
			PoolCache[i].Busy = false
			PoolCache[i].ConversationID = ""
			PoolCache[i].ParentMessageID = ""
			PoolCache[i].User = ""
			fmt.Println("release broker", id)
		}
	}
}
