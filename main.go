package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/youtube_dl_server/src"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	config := src.NewConfig()

	updateNgrok(config)

	r := mux.NewRouter()
	r.HandleFunc("/audio/{url}", audioHandler).Methods("GET")
	http.Handle("/", r)
	log.Println("Starting " + src.MyIp() + ":" + config.NgrokPort)
	log.Fatal(http.ListenAndServe(":"+config.NgrokPort, nil))
}

func updateNgrok(config *src.Config) {
	n := src.NgrokInit(config.NgrokPort, config.NgrokToken)
	f := src.FirebaseServer{
		Ctx:            context.Background(),
		CredentialPath: config.FirebaseTokenPath,
	}
	f.Init()
	f.UpdateData(n)
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
