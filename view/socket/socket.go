package socket

import (
	"github.com/gorilla/websocket"
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

func InitWebSocket() {
	go func() {
		http.HandleFunc("/socket", handler)
		log.Fatal(http.ListenAndServe(":9999", nil))
	}()
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Hello")
	err = conn.WriteJSON("Hello")
	if err != nil {
		log.Println(err)
		return
	}

}
