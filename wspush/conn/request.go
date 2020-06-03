package conn

import (
	"encoding/json"
)

func init() {
	register(new(SubscribeReq))
	register(new(UnsubscribeReq))
	register(new(LoginReq))
	register(new(LogoutReq))
	register(new(Heartbeat))
}

type reqHandler func(conn *Connection, data []byte) error

var msgHandlers = map[string]reqHandler{}

func register(h Handler) {
	msgHandlers[h.Type()] = GenHandler(h)
}

type Handler interface {
	Type() string
	Unmarshal(bs []byte) (Handler, error)
	Handle(c *Connection) error
}

func GenHandler(h Handler) reqHandler {
	return func(c *Connection, bs []byte) error {
		n, err := h.Unmarshal(bs)
		if err != nil {
			return err
		}

		return n.Handle(c)
	}
}

type Request struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
