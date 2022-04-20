package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/youtube-dl-server/core"
)

const (
	TypeConfig = iota
	TypeLog
	TypeState
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

	err = conn.WriteJSON(
		Response{
			TypeIndex: 0,
			Data:      c.LoadConfig(),
		},
	)
	if err != nil {
		log.Println(err)
		return
	}

}

func requestHandler(conn *websocket.Conn) {
	_, data, err := conn.ReadMessage()
	if err != nil {
		log.Println("Bad Request" + string(data))
		log.Println(err)
	}
	log.Println(string(data))
}

func configMessage() ([]byte, error) {
	res := &Response{
		TypeIndex: TypeConfig,
		Data:      c.LoadConfig(),
	}

	return json.Marshal(res)
}
