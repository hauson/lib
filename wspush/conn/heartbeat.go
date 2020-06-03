package conn

import "encoding/json"

type Heartbeat struct{}

func (h *Heartbeat) Unmarshal(data []byte) (Handler, error) {
	return &Heartbeat{}, nil
}

func (h *Heartbeat) Handle(c *Connection) error {
	c.heartbeat()
	return nil
}

func (h *Heartbeat) BelongTo(account string) bool {
	return true
}

func (h *Heartbeat) Type() string {
	return "heartbeat"
}

func (h *Heartbeat) Topic() string {
	return "heartbeat"
}

func (h *Heartbeat) Marshal() json.RawMessage {
	bs, _ := json.Marshal(h)
	return bs
}
