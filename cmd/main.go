/*
Create time at 2023年3月9日0009下午 16:21:17
Create User at Administrator
*/

package main

import (
	"github.com/Skysolderone/BookManagementSys-Golang/pkg/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteredRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
