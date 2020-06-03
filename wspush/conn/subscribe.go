package conn

import "encoding/json"

type SubscribeReq struct {
	Topics []string `json:"topics"`
}

func (req *SubscribeReq) Type() string {
	return "subscribe"
}

func (req *SubscribeReq) Unmarshal(data []byte) (Handler, error) {
	h := &SubscribeReq{}
	if err := json.Unmarshal(data, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (req *SubscribeReq) Handle(c *Connection) error {
	err := c.subscribe(req.Topics)
	c.resp(err)
	return nil
}
