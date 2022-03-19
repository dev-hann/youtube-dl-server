package view

import (
	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/config"
	"net/http"
)

func InitView(r *mux.Router, config *config.ViewConfig) {
	fs := http.FileServer(http.Dir(config.Path))
	r.PathPrefix("/").Handler(fs)
}
