package socket

import (
	"github.com/gorilla/websocket"
	"github.com/youtube-dl-server/core"
	"log"
	"net/http"
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
	log.Println("HEllo")
	conn.WriteMessage(0, []byte("Hello"))
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
