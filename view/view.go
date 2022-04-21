package view

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/youtube-dl-server/config"
)

func InitView(r *mux.Router, config *config.ViewConfig) {
	fs := http.FileServer(http.Dir(config.Path))
	r.PathPrefix("/").Handler(fs)
}
