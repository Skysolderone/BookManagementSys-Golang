/*
Create time at 2023年3月9日0009下午 16:23:03
Create User at Administrator
*/

package routes

import (
	"github.com/Skysolderone/BookManagementSys-Golang/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisteredRoutes = func(r *mux.Router) {
	r.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	r.HandleFunc("/book/{id}", controllers.GetBookbyBookId).Methods("GET")
	r.HandleFunc("/book/{id}", controllers.UpdateBookByBookId).Methods("PUT")
	r.HandleFunc("/book/{id}", controllers.DeleteBookByBookId).Methods("DELETE")
	r.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
}
