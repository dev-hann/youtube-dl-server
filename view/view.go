package view

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Config struct {
	Path string
}

func InitView(r *mux.Router, config *Config) {
	fs := http.FileServer(http.Dir("./view/web/"))
	r.PathPrefix("/").Handler(fs)
}
