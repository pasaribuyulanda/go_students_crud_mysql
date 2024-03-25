package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pasaribuyulanda/go_students_crud_mysql/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterStudentsRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:3307")
	log.Fatal(http.ListenAndServer("localhost:3307", r))
}