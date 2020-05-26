package wspush

import (
	"encoding/json"
	"strings"
)

const (
	topicSeparator string = ":"
	// HeartbeatType heart beat
	HeartbeatType Type = "heartbeat"
	// LogoutType conn logout
	LogoutType Type = "logout"
	// ResponseType response
	ResponseType Type = "response"
	// BalanceType account asset balance
	BalanceType Type = "balance"
	// UserOrderStatusType mov order status
	UserOrderStatusType Type = "order_status"
	// ChainStatusType msg of block height
	ChainStatusType Type = "chain_status"
)

// Type msg type
type Type string

// String to string
func (t Type) String() string {
	return string(t)
}

// Topic return topic string
func (t Type) Topic(ss ...string) string {
	ss = append([]string{t.String()}, ss...)
	return strings.Join(ss, topicSeparator)
}

// WSMsg transmitted  between clint and server
type WSMsg struct {
	Property  *Property       `json:"-"`
	Topic     string          `json:"-"`
	Type      Type            `json:"type"`
	Timestamp uint64          `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
}

// Property msg property
type Property struct {
	// Addresses  msg's belongs to Addresses
	Addresses []string
	Broadcast bool
}

// ChainStatus chain status include Height and so on
type ChainStatus struct {
	Height uint64 `json:"height"`
}

var defaultSubscribeTopics = map[string]bool{
	HeartbeatType.Topic():       true,
	LogoutType.Topic():          true,
	ResponseType.Topic():        true,
	BalanceType.Topic():         true,
	UserOrderStatusType.Topic(): true,
	ChainStatusType.Topic():     true,
}

// IsDefaultSubscribeTopic judge topic is default subscribe
func IsDefaultSubscribeTopic(topic string) bool {
	return defaultSubscribeTopics[topic]
}
