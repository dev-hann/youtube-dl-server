package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"github.com/youtube-dl-server/response"
	"log"
	"net/http"
)

type Api struct {
	config *config.ApiConfig
	core   *core.Core
}

var api *Api

func InitApiHandler(r *mux.Router, config *config.ApiConfig, core *core.Core) {
	initApi(config, core)
	log.Println("/" + config.Version + config.AudioApi)
	r.HandleFunc("/"+config.Version+config.ConfigApi, configHandler).Methods("GET")
	r.HandleFunc("/"+config.Version+config.AudioApi, audioHandler).Methods("GET")
}

func initApi(config *config.ApiConfig, core *core.Core) {
	api = &Api{
		config: config,
		core:   core,
	}
}

func configHandler(writer http.ResponseWriter, request *http.Request) {
	data := api.core.LoadConfig()
	res, err := json.Marshal(response.SuccessResponse(data))
	if err != nil {
		res, _ = json.Marshal(response.FailResponse(err))
	}
	fmt.Fprint(writer, string(res))

}

func audioHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	url := vars["videoID"]
	dlData, err := api.core.LoadAudioURL(url)
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
