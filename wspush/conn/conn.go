package conn

import (
	"fmt"
	"time"
	"encoding/json"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/lib/wspush/msg"
)

const (
	heartbeatTimeout = 2 * time.Minute
)

// Connection wrap web socket
type Connection struct {
	conn     *websocket.Conn
	topics   map[string]bool
	msgQueue chan msg.Raw
	cache    *msg.Cache
	account  string
	exitSig  chan int
}

func New(rawConn *websocket.Conn, cache *msg.Cache) *Connection {
	c := &Connection{
		conn:    rawConn,
		exitSig: make(chan int),
		topics: map[string]bool{
			(*Response)(nil).Type():  true,
			(*Heartbeat)(nil).Type(): true,
		},
		cache:    cache,
		msgQueue: make(chan msg.Raw, 1000),
	}

	go c.readLoop()
	go c.writeLoop()
	return c
}

// IPAddr ip addr
func (c *Connection) IPAddr() string {
	return c.conn.RemoteAddr().String()
}

// Send send msg
func (c *Connection) Send(raw msg.Raw) {
	if raw.BelongTo(c.account) {
		c.write(raw)
	}
}

// IsClosed is closed
func (c *Connection) IsClosed() bool {
	return c.isClose()
}

func (c *Connection) readLoop() {
	for {
		select {
		case <-c.exitSig:
			return
		default:
			c.conn.SetReadDeadline(time.Now().Add(heartbeatTimeout))
			_, bytes, err := c.conn.ReadMessage()
			if err != nil {
				log.WithField("err", err).Error("fail to read message")
				return
			}

			if err := c.handle(bytes); err != nil {
				log.WithField("err", err).Errorf("fail to handle req")
				continue
			}
		}
	}
}

func (c *Connection) writeLoop() {
	for raw := range c.msgQueue {
		if c.isClose() {
			return
		}

		if err := c.conn.WriteJSON(msg.Wrapper(raw)); err != nil {
			log.WithField("err", err).Error("fail to write ws message")
		}
	}
}

func (c *Connection) subscribe(topics []string) error {
	if !c.isLogin() {
		return ErrNotLogin
	}

	for _, topic := range topics {
		c.topics[topic] = true
		if latest, ok := c.cache.Get(topic); ok {
			c.write(latest)
		}
	}

	return nil
}

func (c *Connection) unSubscribe(topics []string) error {
	if !c.isLogin() {
		return ErrNotLogin
	}

	for _, topic := range topics {
		c.topics[topic] = false
	}

	return nil
}

func (c *Connection) heartbeat() {
	c.write(&Heartbeat{})
}

func (c *Connection) login(account string) error {
	if c.isClose() {
		return ErrClosed
	}

	if c.isLogin() && c.account != account {
		return ErrLogined
	}

	c.account = account
	return nil
}

func (c *Connection) logout(account string) error {
	if !c.isLogin() {
		return ErrNotLogin
	}

	if c.account != account {
		return ErrAcount
	}

	c.account = ""
	return nil
}

func (c *Connection) isLogin() bool {
	return c.account != ""
}

func (c *Connection) resp(err error) {
	c.write(newResp(err))
}

func (c *Connection) write(msg msg.Raw) {
	if c.isClose() {
		return
	}

	if c.topics[msg.Topic()] {
		select {
		case c.msgQueue <- msg:
		default:
			log.Warnf("msg remote addr:%s is overflow", c.conn.RemoteAddr().String())
		}
	}
}

func (c *Connection) close() {
	if c.isClose() {
		return
	}

	close(c.exitSig)
	c.conn.Close()
}

func (c *Connection) isClose() bool {
	select {
	case <-c.exitSig:
		return true
	default:
		return false
	}
}

func (c *Connection) handle(bytes []byte) error {
	req := &Request{}
	if err := json.Unmarshal(bytes, req); err != nil {
		return err
	}

	handle, ok := msgHandlers[req.Type]
	if !ok {
		return fmt.Errorf("%s not register handler", req.Type)
	}

	return handle(c, req.Data)
}
