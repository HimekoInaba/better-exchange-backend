package controller

import (
	"grpc-rest-api/src/api/dao"
	"grpc-rest-api/src/api/model"
	"grpc-rest-api/src/api/service"
	"grpc-rest-api/src/api/util"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var (
	userDao = dao.UserDAO{}
)

func RegisterEndpoint(w http.ResponseWriter, r *http.Request)  {
	log.Print("register endpoint hit")
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		log.Println(err)
		return
	}
	user.Id = bson.NewObjectId()

	err := service.Register(user)
	if err != nil {
		log.Print("Error during registration", err)
	}
	util.RespondWithJson(w, http.StatusCreated, user)
}

func GetSingleUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	user, err := userDao.FindById(id)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	users, err := userDao.FindAll()
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(users)
}

func UsersController(router *mux.Router) {
	router.HandleFunc("/register", RegisterEndpoint).Methods("POST")
	router.HandleFunc("/all", GetAllUsers).Methods("GET")
	router.HandleFunc("/{id}", GetSingleUser).Methods("GET")
}