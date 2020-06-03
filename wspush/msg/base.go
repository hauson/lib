package msg

import (
	"encoding/json"
	"time"
)

// Msg transmitted  between clint and server
type Msg struct {
	Type      string          `json:"type"`
	Timestamp int64           `json:"timestamp"`
	Data      json.RawMessage `json:"data"`
	Raw                       `json:"-"`
}

func Wrapper(raw Raw) *Msg {
	return &Msg{
		Type:      raw.Type(),
		Timestamp: time.Now().Unix(),
		Data:      raw.Marshal(),
		Raw:       raw,
	}
}
