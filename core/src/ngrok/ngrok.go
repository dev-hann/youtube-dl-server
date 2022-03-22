package ngrok

import (
	"encoding/json"
	"github.com/youtube-dl-server/config"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Ngrok struct {
	version string
	config  *config.NgrokConfig
	Tunnels []Tunnel `json:"tunnels"`
	Uri     string   `json:"uri"`
}
type Tunnel struct {
	Name      string `json:"name"`
	Uri       string `json:"uri"`
	PublicUrl string `json:"public_url"`
	Proto     string `json:"proto"`
}

func NewNgrok(config *config.NgrokConfig) *Ngrok {
	go ngrokRunCmd(config.Port, config.Token)
	tryCount := 0
	n := Ngrok{
		version: loadVersion(),
		config:  config,
	}
	for len(n.Tunnels) == 0 {
		time.Sleep(1 * time.Second)
		tryCount++
		log.Println("init Ngrok Server => try Count : " + strconv.Itoa(tryCount) + " times..")
		n = *ngrok()
	}
	log.Println("Completed Run Ngrok")
	return &n
}

func loadVersion() string {
	cmd := exec.Command("ngrok", "--version")
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "Version Error"
	}
	return strings.Trim(strings.Split(string(data), " ")[2], "\n")
}

func ngrok() *Ngrok {
	//var res *http.Response
	res, err := http.Get("http://localhost:4040/api/tunnels")
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
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
