package msg

import "encoding/json"

type Raw interface {
	BelongTo(account string) bool
	Type() string
	Topic() string
	Marshal() json.RawMessage
}
