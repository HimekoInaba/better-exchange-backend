package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func RespondWithJson(w http.ResponseWriter, status int, content interface{}) {
	payload, err := json.Marshal(content)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status-Code", string(status))
	w.Write(payload)
}

func RespondWithError(w http.ResponseWriter, status int, msg string) {
	payload, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Status-Code", string(status))
	json.NewEncoder(w).Encode(payload)
}