package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage endpoint hit")
}

func HandleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomePage)

	UsersController(router.PathPrefix("/users").Subrouter())

	handler := cors.Default().Handler(router)
	if err := http.ListenAndServe(":8091", handler); err != nil {
		log.Fatal(err)
	}
}
