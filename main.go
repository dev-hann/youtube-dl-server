package main

import (
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/youtube-dl-server/api"
	"github.com/youtube-dl-server/command"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"github.com/youtube-dl-server/logger"
	"github.com/youtube-dl-server/view"
	"github.com/youtube-dl-server/view/socket"
	"net/http"
)

func main() {
	cmd := command.InitCommand()
	cmd.Parse()
	c := config.NewConfig("./config.yaml")
	logger.InitLogger(c.LoggerConfig)
	appCore := core.InitCore(c)
	r := mux.NewRouter()
	api.InitApiHandler(r, c.ApiConfig, appCore)
	view.InitView(r, c.ViewConfig)
	http.Handle("/", r)
	socket.InitWebSocket(appCore)
	log.Fatal(http.ListenAndServe(":"+c.ApiConfig.Port, nil))

}
