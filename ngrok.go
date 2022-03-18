package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
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

func NgrokInit(port string, token string) *Ngrok {
	go ngrokRunCmd(port, token)
	tryCount := 0
	n := Ngrok{}

	for len(n.Tunnels) == 0 {
		time.Sleep(1 * time.Second)
		tryCount++
		log.Println("init Ngrok Server => try Count : " + strconv.Itoa(tryCount) + " times..")
		n = *ngrok()
	}
	log.Println("Completed Run Ngrok")
	return &n
}

func ngrok() *Ngrok {
	var res *http.Response
	res, err = http.Get("http://localhost:4040/api/tunnels")
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

func ngrokRunCmd(port string, token string) {
	cmd := exec.Command("ngrok", "http", port, "--authtoken", token)
	_, err = cmd.CombinedOutput()
	checkErr()
}
