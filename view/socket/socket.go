package socket

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/youtube-dl-server/core"
	errors "github.com/youtube-dl-server/err"
)

type Socket struct {
	conn *websocket.Conn
	core *core.Core
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func InitWebSocket(core *core.Core) *Socket {
	s := newSocket(core)
	go s.initHandler()
	return s
}

func newSocket(c *core.Core) *Socket {
	s := &Socket{
		core: c,
	}
	return s
}

func (s *Socket) initHandler() {
	http.HandleFunc("/socket", s.handler)
	log.Fatal(http.ListenAndServe(":8888", nil))
}

func (s *Socket) handler(w http.ResponseWriter, r *http.Request) {
	conn, err := s.initConnection(w, r)
	if err != nil {
		log.Println(err)
	}
	go s.initRequestHandler(conn)
}

func (s *Socket) initConnection(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	s.conn = conn
	return conn, nil
}

func (s *Socket) initRequestHandler(conn *websocket.Conn) {
	for true {
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
			err := conn.WriteJSON(s.configMessage())
			if err != nil {
				log.Println(err)
			}
			break
		case TypeLog:
			err := conn.WriteJSON(s.logMessage())
			if err != nil {
				log.Println(err)
			}
			break
		case TypeState:
			err := conn.WriteJSON(s.stateMessage())
			if err != nil {
				log.Println(err)
			}
			break
		}
	}

}
func (s *Socket) configMessage() *Response {
	return &Response{
		Type: TypeConfig,
		Data: s.core.LoadConfig(),
	}
}

func (s *Socket) logMessage() *Response {
	return &Response{
		Type: TypeLog,
		Data: "",
	}
}

func (s *Socket) stateMessage() *Response {
	return &Response{
		Type: TypeState,
		Data: "",
	}
}

func (s *Socket) NewMessage(messsage interface{}) error {
	err := s.conn.WriteJSON(messsage)
	if err != nil {
		return err
	}
	return nil
}
