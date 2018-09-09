package controller

import (
	"fmt"
	"github.com/gorilla/mux"
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

	if err := http.ListenAndServe(":8091", router); err != nil {
		log.Fatal(err)
	}
}
