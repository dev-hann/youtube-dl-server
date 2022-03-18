package main

import (
	"encoding/json"
	"log"
	"time"
)

type Response struct {
	DateTime int64  `json:"dateTime"`
	Result   bool   `json:"result"`
	Data     string `json:"data,omitempty"`
}

func FailResponse(err string) *Response {
	//return marshall(Response{DateTime: time.Now().UnixNano() / int64(time.Millisecond), Result: false, Data: err})
	return &Response{DateTime: time.Now().UnixNano() / int64(time.Millisecond), Result: false, Data: err}
}

func SuccessResponse(data string) *Response {
	//return marshall(Response{DateTime: time.Now().UnixNano() / int64(time.Millisecond), Result: true, Data: data})
	return &Response{DateTime: time.Now().UnixNano() / int64(time.Millisecond), Result: true, Data: data}
}

func marshall(res Response) string {
	resString, err := json.Marshal(res)
	if err != nil {
		log.Println(err)
	}
	return string(resString)
}
