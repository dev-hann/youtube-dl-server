package socket

const (
	TypeConfig = iota
	TypeLog
	TypeState
)

type Request struct {
	Type int `json:"type"`
}

func ConfigRequest() *Request {
	return &Request{
		Type: TypeConfig,
	}
}
