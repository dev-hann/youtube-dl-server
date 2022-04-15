package main

import (
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/api"
	"github.com/youtube-dl-server/argument"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"github.com/youtube-dl-server/logger"
	"github.com/youtube-dl-server/veriosn"
	"github.com/youtube-dl-server/view"
	"github.com/youtube-dl-server/view/socket"
	"log"
	"net/http"
)

func main() {

	arg := argument.InitArgument()
	err := arg.Parse()
	if err != nil {
		log.Print(err)
		return
	}

	arg.Run(
		startServer,
		upgradeServer,
		versionServer,
	)
}

func startServer(configPath string, console *argument.Console) {
	c := config.NewConfig(configPath)
	logger.InitLogger(c.LoggerConfig)
	appCore := core.InitCore(c)
	r := mux.NewRouter()
	api.InitApiHandler(r, c.ApiConfig, appCore)
	view.InitView(r, c.ViewConfig)
	http.Handle("/", r)
	socket.InitWebSocket(appCore)
	log.Fatal(http.ListenAndServe(":"+c.ApiConfig.Port, nil))
}

func upgradeServer(console *argument.Console) {
	v := veriosn.InitVersion()
	var res []byte
	needUpgrade := true
	for needUpgrade {
		res, needUpgrade = v.CheckVersion()
		if res != nil {
			console.Log(string(res))
		}
		if needUpgrade {
			console.ShowLogo()
			res, err := v.PullNewVersion()
			if err != nil {
				console.Log(string(res))
				return
			}
			res, err = v.Build()
			if err != nil {
				console.Log(string(res))
				return
			}
		} else {
			console.Log("Current Version is already Newest. " + v.CurrentVersion())
		}
	}

}

func versionServer(console *argument.Console) {
	v := veriosn.InitVersion()
	console.ShowLogo()
	console.Log(v.CurrentVersion())
}
