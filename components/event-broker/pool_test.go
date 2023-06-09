package EventBroker_test

import (
	"testing"

	eb "github.com/soulteary/sparrow/components/event-broker"
)

func TestNewBrokerManager(t *testing.T) {
	size := 5
	bp := eb.NewBrokerManager(size)
	if bp == nil {
		t.Errorf("Returned BrokersPool should not be nil")
	}
	if len(eb.PoolCache) != size {
		t.Errorf("Expected PoolCache length of %d, but got %d", size, len(eb.PoolCache))
	}
}

func TestGetBroker(t *testing.T) {
	// Test when there are available brokers
	bp := eb.NewBrokerManager(2)
	broker1 := bp.GetBroker("user1", "conv1", "")
	if broker1 == nil {
		t.Errorf("Returned Broker should not be nil")
	}
	if !eb.PoolCache[0].Busy {
		t.Errorf("Expected pool 0 to be busy")
	}
	if eb.PoolCache[0].User != "user1" || eb.PoolCache[0].ConversationID != "conv1" ||
		eb.PoolCache[0].ParentMessageID != "" {
		t.Errorf("Unexpected value for pool 0: %+v", eb.PoolCache[0])
	}

	// Test when all brokers are busy and there is no matching parent message ID or conversation ID
	broker2 := bp.GetBroker("user2", "", "")
	if broker2 == nil {
		t.Errorf("Returned Broker should not be nil")
	}
	if !eb.PoolCache[1].Busy {
		t.Errorf("Expected pool 1 to be busy")
	}
	if eb.PoolCache[1].User != "user2" || eb.PoolCache[1].ConversationID != "" ||
		eb.PoolCache[1].ParentMessageID != "" {
		t.Errorf("Unexpected value for pool 1: %+v", eb.PoolCache[1])
	}

	// Test when all brokers are busy and there is a matching parent message ID
	broker3 := bp.GetBroker("user3", "", "pm1")
	if broker3 == nil {
		t.Errorf("Returned Broker should not be nil")
	}
	if !eb.PoolCache[0].Busy {
		t.Errorf("Expected pool 0 to be busy")
	}

	// TODO: Fix this test

	// if eb.PoolCache[0].User != "user3" || eb.PoolCache[0].ConversationID != "" ||
	// 	eb.PoolCache[0].ParentMessageID != "pm1" {
	// 	t.Errorf("Unexpected value for pool 0: %+v", eb.PoolCache[0])
	// }

	// Test when all brokers are busy and there is a matching conversation ID
	// broker4 := bp.GetBroker("user4", "conv2", "")
	// if broker4 == nil {
	// 	t.Errorf("Returned Broker should not be nil")
	// }
	// if !eb.PoolCache[1].Busy {
	// 	t.Errorf("Expected pool 1 to be busy")
	// }
	// if eb.PoolCache[1].User != "user4" || eb.PoolCache[1].ConversationID != "conv2" ||
	// 	eb.PoolCache[1].ParentMessageID != "" {
	// 	t.Errorf("Unexpected value for pool 1: %+v", eb.PoolCache[1])
	// }

	// Test when there are no available brokers
	// broker5 := bp.GetBroker("user5", "", "")
	// if broker5 != nil {
	// 	t.Errorf("Expected returned Broker to be nil")
	// }
}

func TestFreePool(t *testing.T) {
	bp := eb.NewBrokerManager(1)
	bp.GetBroker("user1", "conv1", "")
	eb.FreePool(0)
	if eb.PoolCache[0].Busy {
		t.Errorf("Expected pool 0 to be free")
	}
	if eb.PoolCache[0].User != "" || eb.PoolCache[0].ConversationID != "" ||
		eb.PoolCache[0].ParentMessageID != "" {
		t.Errorf("Unexpected value for pool 0: %+v", eb.PoolCache[0])
	}
}
