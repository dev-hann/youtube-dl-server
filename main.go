package main

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	config := NewConfig()

	updateNgrok(config)

	r := mux.NewRouter()
	r.HandleFunc("/audio/{url}", audioHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+config.NgrokPort, nil))
}

func updateNgrok(config *Config) {
	n := NgrokInit(config.NgrokPort, config.NgrokToken)
	f := FirebaseServer{
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
	var res *Response
	if err != nil {
		res = FailResponse(dlURL)
	} else {
		res = SuccessResponse(dlURL)
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
