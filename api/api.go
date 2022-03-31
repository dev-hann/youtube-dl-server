package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"net/http"
)

type Api struct {
	config *config.ApiConfig
	core   *core.Core
}

var api *Api
var err error
var logger *log.Entry

func initLogger() {
	logger = log.WithFields(
		log.Fields{
			"Field": "Api",
		},
	)
}

func callApi(request *http.Request) {
	logger.WithFields(
		log.Fields{
			"userAgent": request.UserAgent(),
		},
	).Info(request.URL.String())
}
func checkError() {
	if err != nil {
		logger.Error(err)
	}
	err = nil
}

func initApi(config *config.ApiConfig, core *core.Core) {
	initLogger()
	api = &Api{
		config: config,
		core:   core,
	}
}

func InitApiHandler(r *mux.Router, config *config.ApiConfig, core *core.Core) {
	initApi(config, core)
	r.HandleFunc("/"+config.Version+config.ConfigApi, configHandler).Methods("GET")
	r.HandleFunc("/"+config.Version+config.AudioApi, audioHandler).Methods("GET")
	r.HandleFunc("/"+config.Version+config.MelonApi, melonHandler).Methods("GET")
	r.HandleFunc("/"+config.Version+config.YoutubeApi, youtubeChartHandler).Methods("GET")
}

func youtubeChartHandler(writer http.ResponseWriter, request *http.Request) {
	callApi(request)
	var res interface{}
	res, err = api.core.LoadYoutubeChart()
	checkError()
	responseData(writer, res)
}

func melonHandler(writer http.ResponseWriter, request *http.Request) {
	callApi(request)
	var res interface{}
	res, err = api.core.LoadMelonChart()
	checkError()
	responseData(writer, res)
}

func configHandler(writer http.ResponseWriter, request *http.Request) {
	callApi(request)
	res := api.core.LoadConfig()
	responseData(writer, res)
}

func audioHandler(writer http.ResponseWriter, request *http.Request) {
	callApi(request)
	vars := mux.Vars(request)
	url := vars["videoID"]
	var res []byte
	res, err = api.core.LoadAudioURL(url)
	checkError()
	responseData(writer, res)

	//dlURL := string(dlData)
	//fmt.Println(dlURL)
	//var res *Response
	//if err != nil {
	//	res = FailResponse(dlURL)
	//} else {
	//	res = SuccessResponse(dlURL)
	//}
	//e := json.NewEncoder(writer)
	//e.SetEscapeHTML(false)
	//e.Encode(res)
}

func responseData(writer http.ResponseWriter, response interface{}) {
	var res []byte
	res, err = json.Marshal(SuccessResponse(response))
	checkError()
	if err != nil {
		res, err = json.Marshal(FailResponse(err))
		checkError()
	}
	_, err = fmt.Fprint(writer, string(res))
	checkError()

}
