package main

import (
	"github.com/Abdelwhab-1/book_store_user_service-/app"
	"github.com/gorilla/mux"
)
var router *mux.Router
func init(){
	router = mux.NewRouter()
}
func main() {
	app.StartApp(router)
}
