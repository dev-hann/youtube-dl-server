package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/youtube_dl_server/src"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {

	port := os.Getenv("port")
	if port == "" {
		port = "8444"
	}

	r := mux.NewRouter()
	r.HandleFunc("/audio/{url}", audioHandler).Methods("GET")
	http.Handle("/", r)
	log.Println("Starting " + src.MyIp() + ":" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func audioHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	url := vars["url"]
	dlURL, err := loadAudioURL(url)
	var res *src.Response
	if err != nil {
		res = src.FailResponse(dlURL)
	} else {
		res = src.SuccessResponse(dlURL)
	}
	e := json.NewEncoder(writer)
	e.SetEscapeHTML(false)
	e.Encode(res)
}

func loadAudioURL(url string) (string, error) {
	log.Println("request audio url => " + url)
	cmd := exec.Command("youtube-dl", "-x", "--audio-format", "mp3", url, "--get-url")
	out, err := cmd.CombinedOutput()
	//log.Println(string(out))
	return string(out), err
}
