package socket

type Response struct {
	TypeIndex int         `json:"type_index"`
	Data      interface{} `json:"data"`
	Error     string      `json:"error"`
}
