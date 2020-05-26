package wspush

import (
	"encoding/json"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/bytom/blockcenter/wspush/msg"
)

const (
	heartbeatTimeout = 2 * time.Minute
)

// connection wrap web socket
type connection struct {
	loggedAccount *account
	conn          *websocket.Conn
	server        *Server
	topics        *sync.Map
	exitSig       chan int
	msgQueue      chan *msg.WSMsg
	mutex         sync.RWMutex
}

func newConn(rawConn *websocket.Conn, server *Server, account *account, topics *sync.Map) *connection {
	return &connection{
		conn:          rawConn,
		server:        server,
		loggedAccount: account,
		exitSig:       make(chan int),
		topics:        topics,
		msgQueue:      make(chan *msg.WSMsg, 1000),
	}
}

func (c *connection) logout() {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.close()
}

func (c *connection) subscribe(topics []string) error {
	for _, topic := range topics {
		c.topics.Store(topic, nil)
	}
	return nil
}

func (c *connection) unSubscribe(topics []string) {
	for _, topic := range topics {
		c.topics.Delete(topic)
	}
}

func (c *connection) start() {
	go c.readMsgLoop()
	go c.writeMsgLoop()
}

func (c *connection) send(msg *msg.WSMsg) {
	if c.isSubscribe(msg.Topic) {
		select {
		case c.msgQueue <- msg:
		default:
			log.Warnf("msg queue of wallet of address:%s is overflow", c.loggedAccount.address)
		}
	}
}

func (c *connection) write(msg *msg.WSMsg) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	if c.isClose() {
		return
	}

	if err := c.conn.WriteJSON(msg); err != nil {
		log.WithField("err", err).Error("fail to write ws message")
	}
}

func (c *connection) close() {
	if c.isClose() {
		return
	}

	close(c.exitSig)
	_ = c.conn.Close()
}

func (c *connection) isClose() bool {
	select {
	case <-c.exitSig:
		return true
	default:
		return false
	}
}

func (c *connection) isSubscribe(topic string) bool {
	if msg.IsDefaultSubscribeTopic(topic) {
		return true
	}

	_, ok := c.topics.Load(topic)
	return ok
}

func (c *connection) readMsgLoop() {
	for {
		select {
		case <-c.exitSig:
			return
		default:
			_ = c.conn.SetReadDeadline(time.Now().Add(heartbeatTimeout))

			_, bytes, err := c.conn.ReadMessage()
			if err != nil {
				log.WithField("err", err).Error("fail to read message")

				req := &request{Type: msgLogout}
				handleMsg(req, c)
				return
			}

			req := &request{}
			if err := json.Unmarshal(bytes, req); err != nil {
				log.WithField("err", err).Errorf("fail to json unmarshal: %v", req)
				continue
			}

			handleMsg(req, c)
		}
	}
}

func (c *connection) writeMsgLoop() {
	for msg := range c.msgQueue {
		c.write(msg)
	}
}
