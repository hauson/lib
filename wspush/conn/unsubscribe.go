package conn

import "encoding/json"

type UnsubscribeReq struct {
	Topics []string `json:"topics"`
}

func (req *UnsubscribeReq) Type() string {
	return "unsubscribe"
}

func (req *UnsubscribeReq) Unmarshal(data []byte) (Handler, error) {
	h := &UnsubscribeReq{}
	if err := json.Unmarshal(data, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (req *UnsubscribeReq) Handle(c *Connection) error {
	err := c.unSubscribe(req.Topics)
	c.resp(err)
	return nil
}
