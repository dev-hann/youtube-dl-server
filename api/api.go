package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"net/http"
)

type Api struct {
	config *config.ApiConfig
	core   *core.Core
}

var api *Api

func InitApiHandler(r *mux.Router, config *config.ApiConfig, core *core.Core) {
	initApi(config, core)
	r.HandleFunc("/"+config.Version+config.ConfigApi, configHandler).Methods("GET")
	r.HandleFunc("/"+config.Version+config.AudioApi, audioHandler).Methods("GET")
	r.HandleFunc("/"+config.Version+config.MelonApi, melonHandler).Methods("GET")
}

func initApi(config *config.ApiConfig, core *core.Core) {
	api = &Api{
		config: config,
		core:   core,
	}
}

func melonHandler(writer http.ResponseWriter, request *http.Request) {
	m := api.core.LoadMelonChart()
	res, err := json.Marshal(SuccessResponse(m))
	if err != nil {
		res, _ = json.Marshal(FailResponse(err))
	}
	fmt.Fprint(writer, string(res))

}

func configHandler(writer http.ResponseWriter, request *http.Request) {
	data := api.core.LoadConfig()
	res, err := json.Marshal(SuccessResponse(data))
	if err != nil {
		res, _ = json.Marshal(FailResponse(err))
	}
	fmt.Fprint(writer, string(res))

}

func audioHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	url := vars["videoID"]
	dlData, err := api.core.LoadAudioURL(url)
	dlURL := string(dlData)
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
