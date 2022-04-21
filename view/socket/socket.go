package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/youtube-dl-server/core"
	errors "github.com/youtube-dl-server/err"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var c *core.Core

func InitWebSocket(core *core.Core) {
	c = core
	go func() {
		http.HandleFunc("/socket", handler)
		log.Fatal(http.ListenAndServe(":8888", nil))
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	go requestHandler(conn)

	if err != nil {
		log.Println(err)
		return
	}

}

func requestHandler(conn *websocket.Conn) {
	_, data, err := conn.ReadMessage()
	if err != nil {
		log.Println(errors.BadRequest)
		log.Println(err)
	}
	var req Request
	err = json.Unmarshal([]byte(data), &req)
	if err != nil {
		log.Println(err)
		return
	}
	switch req.Type {
	case TypeConfig:
		err := conn.WriteJSON(configMessage())
		if err != nil {
			log.Println(err)
		}
		break

	}
}

func configMessage() *Response {
	return &Response{
		TypeIndex: TypeConfig,
		Data:      c.LoadConfig(),
	}

}
