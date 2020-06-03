package conn

import (
	"encoding/json"
)

const (
	success = 200
	failure = 0
)

type Response struct {
	ErrMsg  string `json:"err_msg"`
	ErrCode uint64 `json:"err_code"`
}

func newResp(err error) *Response {
	res := &Response{}
	if err == nil {
		res.ErrCode = success
	} else {
		res.ErrCode = failure
		res.ErrMsg = err.Error()
	}
	return res
}

func (res *Response) BelongTo(account string) bool {
	return true
}

func (res *Response) Type() string {
	return "response"
}

func (res *Response) Topic() string {
	return "response"
}

func (res *Response) Marshal() json.RawMessage {
	bs, _ := json.Marshal(res)
	return bs
}
