package conn

import "encoding/json"

type LogoutReq struct {
	Account string `json:"account"`
}

func (req *LogoutReq) Type() string {
	return "logout"
}

func (req *LogoutReq) Unmarshal(data []byte) (Handler, error) {
	h := &LogoutReq{}
	if err := json.Unmarshal(data, h); err != nil {
		return nil, err
	}

	return h, nil
}

func (req *LogoutReq) Handle(c *Connection) error {
	err := c.logout(req.Account)
	c.resp(err)
	return nil
}
