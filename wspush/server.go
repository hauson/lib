package wspush

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"

	"github.com/bytom/blockcenter/wspush/msg"
	"github.com/bytom/blockcenter/config"
)

// Server a conn container
type Server struct {
	*connManager
	rwMutex sync.RWMutex
	cfg     config.Wspush
}

// NewServer return Server
func NewServer(cfg config.Wspush) *Server {
	server := &Server{
		cfg:         cfg,
		connManager: newConnManager(),
	}
	return server
}

type account struct {
	address      string
	platform     string
	deviceNumber string
}

func (a *account) key() string {
	return fmt.Sprintf("%s:%s:%s", a.address, a.platform, a.deviceNumber)
}

// Run used to start the websocket servewr
func (s *Server) Run() {
	router := gin.Default()
	router.GET("/api/v3/vapor/websocket", s.websocketHandler)
	addr := fmt.Sprintf(":%d", s.cfg.GinGonic.ListeningPort)
	log.Fatal(router.Run(addr))
}

// Broadcast used to broadcast the message to all connection
func (s *Server) Broadcast(msg *msg.WSMsg) {
	for _, conn := range s.connManager.all() {
		conn.send(msg)
	}
}

// SendByAddress send the message to specify connection
func (s *Server) SendByAddress(address string, msg *msg.WSMsg) {
	for _, conn := range s.connManager.getByAddress(address) {
		conn.send(msg)
	}
}

func (s *Server) websocketHandler(ctx *gin.Context) {
	upgrader := &websocket.Upgrader{}
	wsConn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.WithField("err", err).Error("fail to create ws connection")
		return
	}

	address := ctx.Query("address")
	if address == "" {
		log.Error("fail to get address from request")
		return
	}

	account := &account{
		address:      address,
		platform:     ctx.Query("platform"),
		deviceNumber: ctx.Query("device_number"),
	}

	topics := new(sync.Map)
	if preConn := s.connManager.getByAccount(account); preConn != nil {
		topics = preConn.topics
		preConn.logout()
	}

	conn := newConn(wsConn, s, account, topics)
	s.connManager.add(conn)
	conn.start()
}
