package client

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"

	"github.com/hauson/lib/wspush/conn"
	"github.com/hauson/lib/wspush/msg"
)

type Client struct {
	url     string
	conn    *websocket.Conn
	exitSig chan int
}

func New(ip, port string) *Client {
	return &Client{
		url:     fmt.Sprintf("ws://%s:%s/websocket", ip, port),
		exitSig: make(chan int),
	}
}

func (c *Client) Connect() error {
	conn, _, err := websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		return err
	}

	c.conn = conn
	return nil
}

func (c *Client) Close() {
	c.Close()
	close(c.exitSig)
}

func (c *Client) ReadLoop() {
	for {
		select {
		case <-c.exitSig:
			return
		default:
			_, bs, err := c.conn.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			message := &msg.Msg{}
			if err := json.Unmarshal(bs, message); err != nil {
				log.Println("unmarshal err:", err)
				return
			}

			if message.Type == "chain_status" {
				chainStatus := &ChainStatus{}
				if err := json.Unmarshal(message.Data, chainStatus); err == nil {
					fmt.Println("chain_status", chainStatus)
				}
			} else {
				fmt.Println(message.Type, message.Timestamp, string(message.Data))
			}
		}
	}
}

func (c *Client) HeartbeatLoop(du time.Duration) {
	ticker := time.NewTicker(du)
	defer ticker.Stop()

	for ; true; <-ticker.C {
		c.conn.WriteJSON(&conn.Request{Type: "heartbeat"})
	}
}

func (c *Client) reqSubscribe(topic string) error {
	sub := conn.SubscribeReq{Topics: []string{topic}}
	bs, _ := json.Marshal(sub)
	return c.conn.WriteJSON(&conn.Request{Type: sub.Type(), Data: bs})
}

func (c *Client) reqLogin(account string) error {
	login := &conn.LoginReq{Account: account}
	bs, _ := json.Marshal(login)
	return c.conn.WriteJSON(&conn.Request{Type: login.Type(), Data: bs})
}
