package services

import (
	"github.com/Abdelwhab-1/book_store_user_service-/app"
	"net/http"
)

func main() {
	srv = http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: app.Mux,
	}
}
