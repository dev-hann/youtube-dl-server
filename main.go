package main

import (
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/api"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/core"
	"github.com/youtube-dl-server/view"
	"log"
	"net/http"
)

func main() {
	c := config.NewConfig("./config.yaml")
	appCore := core.InitCore(c)
	r := mux.NewRouter()
	api.InitApiHandler(r, c.ApiConfig, appCore)
	view.InitView(r, c.ViewConfig)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+c.NgrokConfig.Port, nil))
}
