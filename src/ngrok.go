package src

import (
	"encoding/json"
	"io/ioutil"
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
	res, err = http.Get("http://127.0.0.1:4040/api/tunnels")
	checkErr()
	defer res.Body.Close()
	var body []byte
	body, err = ioutil.ReadAll(res.Body)
	checkErr()
	var n Ngrok
	err = json.Unmarshal(body, &n)
	checkErr()
	return &n
}
