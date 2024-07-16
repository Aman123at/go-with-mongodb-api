package route

import (
	"github.com/Aman123at/usermanage/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", controller.Welcome).Methods("GET")
	router.HandleFunc("/api/user/all", controller.GetAllUsers).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.GetUserById).Methods("GET")
	router.HandleFunc("/api/user/{id}", controller.UpdateUserById).Methods("PUT")
	router.HandleFunc("/api/user/{id}", controller.DeleteUserById).Methods("DELETE")
	router.HandleFunc("/api/user", controller.AddOneUser).Methods("POST")
	return router
}
