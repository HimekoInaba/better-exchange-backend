package controller

import (
	"better-exchange-backend/src/model"
	"better-exchange-backend/src/service"
	"better-exchange-backend/src/util"
	"better-exchange-backend/src/core/database"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

var (
	userDao = database.Connection{}
)

func Register(w http.ResponseWriter, r *http.Request) {
	log.Print("register endpoint hit")
	defer r.Body.Close()
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		log.Println(err)
		return
	}

	oldUser, err := userDao.FindByEmail(user.Email)
	if err != nil {
		log.Println(err)
	}
	if oldUser != (model.User{}) {
		util.RespondWithError(w, http.StatusConflict, "User with this email already registered")
	} else {
		user.Id = bson.NewObjectId()
		err := service.Register(user)
		if err != nil {
			log.Print("Error during registration", err)
		}
		util.RespondWithJson(w, http.StatusCreated, user)
	}
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

func Login(w http.ResponseWriter, r *http.Request) {
	log.Println("Login endpoint hit")
	var loginData model.LoginData

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		util.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		log.Println(err)
		return
	}

	valid, err := service.Login(loginData)
	if err != nil {
		log.Println(err)
	}
	if valid {
		json.NewEncoder(w).Encode("Success")
	} else {
		json.NewEncoder(w).Encode("Fail")
	}
}

func UsersController(router *mux.Router) {
	router.HandleFunc("/register", Register).Methods("POST")
	router.HandleFunc("/all", GetAllUsers).Methods("GET")
	router.HandleFunc("/{id}", GetSingleUser).Methods("GET")
	router.HandleFunc("/login", Login).Methods("POST")
}