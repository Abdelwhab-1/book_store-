package app

import (
	"github.com/gorilla/mux"
	"net/http"
	_ "net/http"
)

func StartApp(mux *mux.Router) {
	mapper(mux)
	srv := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}
	srv.ListenAndServe()
}
