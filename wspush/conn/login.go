package conn

import "encoding/json"

type LoginReq struct {
	Account string `json:"account"`
}

func (req *LoginReq) Type() string {
	return "login"
}

func (req *LoginReq) Unmarshal(data []byte) (Handler, error) {
	h := &LoginReq{}
	if err := json.Unmarshal(data, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (req *LoginReq) Handle(c *Connection) error {
	err := c.login(req.Account)
	c.resp(err)
	return nil
}
