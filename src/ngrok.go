package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Ngrok struct {
	Tunnels []Tunnel `json:"tunnels"`
	Uri     string   `json:"uri"`
}
type Tunnel struct {
	Name      string `json:"name"`
	Uri       string `json:"uri"`
	PublicUrl string `json:"public_url"`
	Proto     string `json:"proto"`
}

func NgrokInit() *Ngrok {
	var res *http.Response
	res, err = http.Get("http://ngrok:4040/api/tunnels")
	//res, err = http.Get("http://172.28.0.2:4040/api/tunnels")
	checkErr()
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	checkErr()
	log.Println(string(body))
	var n Ngrok
	err = json.Unmarshal(body, &n)
	checkErr()
	return &n
}
