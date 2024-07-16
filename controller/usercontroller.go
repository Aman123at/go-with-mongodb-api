package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Aman123at/usermanage/db"
	"github.com/Aman123at/usermanage/model"
	"github.com/gorilla/mux"
)

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to mongo db api backend server</h1>"))
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	allUsers := db.GetAllUsersFromDB()

	json.NewEncoder(w).Encode(allUsers)
}

func AddOneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal(err)
	}

	if user.IsEmpty() {
		json.NewEncoder(w).Encode(map[string]string{"message": "User should not be empty"})
	} else {
		db.InsertOneUserInDB(user)
		json.NewEncoder(w).Encode(map[string]string{"message": "User added successfully"})
	}
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	user := db.GetOneUserFromDB(params["id"])

	if len(user) > 0 {
		json.NewEncoder(w).Encode(user[0])
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "No user found by this id"})
	}
}

func UpdateUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	var jsonBody map[string]string

	json.NewDecoder(r.Body).Decode(&jsonBody)

	db.UpdateOneUserInDB(params["id"], jsonBody)

	json.NewEncoder(w).Encode(map[string]string{"message": "User updated successfully"})
}

func DeleteUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	db.DeleteOneUserFromDB(params["id"])

	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}
