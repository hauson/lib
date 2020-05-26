package wspush

import (
	"encoding/json"

	"github.com/sirupsen/logrus"

	"github.com/bytom/blockcenter/wspush/msg"
)

const (
	msgLogout      = "logout"
	msgSubscribe   = "subscribe"
	msgUnsubscribe = "unsubscribe"
	msgHeartbeat   = "heartbeat"

	success = 200
	failure = 0
)

type msgHandler func(req *request, conn *connection)

var msgHandlers = map[string]msgHandler{
	msgLogout:      logoutHandler,
	msgSubscribe:   subscribeHandler,
	msgUnsubscribe: unSubscribeHandler,
	msgHeartbeat:   heartBeatHandler,
}

type request struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

type response struct {
	Type   string `json:"type"`
	ErrMsg string `json:"err_msg"`
	// ErrCode=200 is success, other is failure
	ErrCode uint64 `json:"err_code"`
}

// HandleMsg handle request with recover
func handleMsg(req *request, conn *connection) {
	defer func() {
		if err := recover(); err != nil {
			logrus.Errorf("Requester Handler caught Exception: %v \n", err)
		}
	}()

	handleFun, ok := msgHandlers[req.Type]
	if !ok {
		logrus.Errorf("%v not register handleMsg", req.Type)
		return
	}

	handleFun(req, conn)
}

func successResp(reqType string, conn *connection) {
	resp := &response{Type: reqType, ErrCode: success}
	sendResp(conn, resp)
}

func errorResp(err error, reqType string, conn *connection) {
	resp := &response{Type: reqType, ErrCode: failure}
	resp.ErrMsg = err.Error()
	sendResp(conn, resp)
}

func sendResp(conn *connection, resp *response) {
	body, _ := json.Marshal(resp)
	conn.send(&msg.WSMsg{
		Type:  msg.ResponseType,
		Topic: msg.ResponseType.Topic(),
		Data:  body,
	})
}

type subscribeRaw struct {
	Topics []string `json:"topics"`
}

func subscribeHandler(req *request, conn *connection) {
	sub := &subscribeRaw{}
	if err := json.Unmarshal(req.Data, sub); err != nil {
		errorResp(err, req.Type, conn)
		return
	}

	err := conn.subscribe(sub.Topics)
	if err != nil {
		errorResp(err, req.Type, conn)
		return
	}

	successResp(req.Type, conn)
}

type unsubscribeRaw struct {
	Topics []string `json:"topics"`
}

func unSubscribeHandler(req *request, conn *connection) {
	unSub := &unsubscribeRaw{}
	if err := json.Unmarshal(req.Data, unSub); err != nil {
		errorResp(err, req.Type, conn)
		return
	}

	conn.unSubscribe(unSub.Topics)
	successResp(req.Type, conn)
	return
}

func heartBeatHandler(_ *request, conn *connection) {
	successResp(string(msg.HeartbeatType), conn)
	return
}

func logoutHandler(req *request, conn *connection) {
	successResp(req.Type, conn)
	conn.logout()
	conn.server.del(conn)
}
