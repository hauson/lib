package pusher

import (
	"net/http"

	"github.com/gorilla/websocket"

	"github.com/lib/wspush/configs"
	"github.com/lib/wspush/conn"
	"github.com/lib/wspush/msg"
)

// Pusher a conn container
type Pusher struct {
	cfg      *configs.Config
	upgrader *websocket.Upgrader
	conns    *connects
	cache    *msg.Cache
}

// Wrapper return Pusher
func New(cfg *configs.Config) *Pusher {
	server := &Pusher{
		cfg:      cfg,
		conns:    newConns(),
		cache:    msg.NewCache(),
		upgrader: new(websocket.Upgrader),
	}
	return server
}

// Run used to start the web socket server
func (s *Pusher) Run() {
	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		c, err := s.upgrader.Upgrade(w, r, nil)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		s.conns.add(conn.New(c, s.cache))
	})

	http.ListenAndServe(s.cfg.Port, nil)
}

// Send msg
func (s *Pusher) Send(msg msg.Raw) {
	s.cache.Set(msg)

	for _, c := range s.conns.all() {
		if c.IsClosed() {
			s.conns.del(c)
		} else {
			c.Send(msg)
		}
	}
}
