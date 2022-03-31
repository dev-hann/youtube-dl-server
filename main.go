package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/youtube-dl-server/api"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"github.com/youtube-dl-server/logger"
	"github.com/youtube-dl-server/view"
	"net/http"
)

func main() {
	c := config.NewConfig("./config.yaml")
	logger.InitLogger(c.LoggerConfig)
	appCore := core.InitCore(c)
	r := mux.NewRouter()
	api.InitApiHandler(r, c.ApiConfig, appCore)
	view.InitView(r, c.ViewConfig)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+c.ApiConfig.Port, nil))
}
