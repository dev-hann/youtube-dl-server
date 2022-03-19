package main

import (
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/api"
	"github.com/youtube-dl-server/config"
	"github.com/youtube-dl-server/firebase"
	"github.com/youtube-dl-server/ngrok"
	"github.com/youtube-dl-server/view"
	"github.com/youtube-dl-server/youtube_dl"
	"log"
	"net/http"
)

func main() {

	c := config.NewConfig("./config.yaml")
	dl := youtube_dl.NewYoutubeDL(c.YoutubeDlConfig)
	n := ngrok.NewNgrok(c.NgrokConfig)
	f := firebase.NewFirebase(c.FirebaseTokenPath)
	f.UpdateNgrok(n)

	r := mux.NewRouter()
	view.InitView(r, c.ViewConfig)
	api.InitApiHandler(r, c.ApiConfig, dl)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+c.NgrokConfig.Port, nil))
}
