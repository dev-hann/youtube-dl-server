package ngrok

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"
)

type Config struct {
	Port  string
	Token string
}

type Ngrok struct {
	config  *Config
	Tunnels []Tunnel `json:"tunnels"`
	Uri     string   `json:"uri"`
}
type Tunnel struct {
	Name      string `json:"name"`
	Uri       string `json:"uri"`
	PublicUrl string `json:"public_url"`
	Proto     string `json:"proto"`
}

func NewNgrok(config *Config) *Ngrok {
	go ngrokRunCmd(config.Port, config.Token)
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
	//var res *http.Response
	res, err := http.Get("http://localhost:4040/api/tunnels")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	//var body []byte
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panicln(err)
	}
	var n Ngrok
	err = json.Unmarshal(body, &n)
	if err != nil {
		log.Panicln(err)
	}
	return &n
}

func ngrokRunCmd(port string, token string) {
	cmd := exec.Command("ngrok", "http", port, "--authtoken", token)
	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Panicln(err)
	}
}
