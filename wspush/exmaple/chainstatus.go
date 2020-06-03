package main

import "encoding/json"

// ChainStatus chain status include Height and so on
type ChainStatus struct {
	Accounts []string `json:"-"`
	Height   uint64   `json:"height"`
}

func (c *ChainStatus) BelongTo(account string) bool {
	for _, msgAccount := range c.Accounts {
		if account == msgAccount {
			return true
		}
	}

	return false
}

func (c *ChainStatus) Type() string {
	return "chain_status"
}

func (c *ChainStatus) Topic() string {
	return "chain_status"
}

func (c *ChainStatus) Marshal() json.RawMessage {
	bs, _ := json.Marshal(c)
	return bs
}
