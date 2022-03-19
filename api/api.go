package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/response"
	"github.com/youtube-dl-server/youtube_dl"
	"net/http"
)

type handleFunc = func(http.ResponseWriter, *http.Request)

var dl *youtube_dl.YoutubeDL

type Config struct {
	Version  string
	AudioAPI string
}

func InitApiHandler(r *mux.Router, config *Config, youtubeDl *youtube_dl.YoutubeDL) {
	dl = youtubeDl
	config.handler(r, config.AudioAPI, audioHandler).Methods("GET")
}

func (c *Config) handler(r *mux.Router, path string, f handleFunc) *mux.Route {
	return r.HandleFunc("/"+c.Version+path, f)
}

func audioHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	url := vars["videoID"]
	dlData, err := dl.LoadAudio(url)
	dlURL := string(dlData)
	var res *response.Response
	if err != nil {
		res = response.FailResponse(dlURL)
	} else {
		res = response.SuccessResponse(dlURL)
	}
	e := json.NewEncoder(writer)
	e.SetEscapeHTML(false)
	e.Encode(res)
}
