package main

import (
	"log"
	"net/http"
	"os"

	controllers "gem-exp/app/controllers"

	"github.com/gorilla/mux"
)

func route() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/", controllers.Hello)
	r.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./"))))
	return r
}

func main() {
	//InitConfig()
	r := route()
	http.Handle("/", r)
	if os.Getenv("MODE") != "" {
		log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))
	} else {
		log.Fatal(http.ListenAndServe(":"+"10000", nil))
	}
}
