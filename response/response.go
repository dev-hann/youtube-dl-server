package response

import (
	"encoding/json"
	"log"
	"time"
)

type Response struct {
	DateTime int64       `json:"dateTime"`
	Result   bool        `json:"result"`
	Data     interface{} `json:"data,omitempty"`
}

func currentTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func FailResponse(err interface{}) *Response {
	return &Response{DateTime: currentTime(), Result: false, Data: err}
}

func SuccessResponse(data interface{}) *Response {
	return &Response{DateTime: currentTime(), Result: true, Data: data}
}

func marshall(res Response) string {
	resString, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	return string(resString)
}
