package src

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"sync"
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

func initCmd() {
	log.Println("init ngrok..")
	log.Println("run ngrok")
	cmd := exec.Command("ngrok", "http", "8444")
	go cmd.CombinedOutput()
	wg.Done()
}

var wg sync.WaitGroup

func NgrokInit() *Ngrok {
	wg.Add(1)
	initCmd()
	wg.Wait()
	log.Println("done ngrok")
	var res *http.Response
	res, err = http.Get("http://localhost:4040/api/tunnels")
	//res, err = http.Get("http://172.28.0.2:4040/api/tunnels")
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
