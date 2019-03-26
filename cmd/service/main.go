package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sasimpson/servicedemo/api"
	"github.com/sasimpson/servicedemo/interfaces/mock"
)

//main will assemble and start our service as well as handle all the configuration bits
func main() {

	h := api.Handler{
		User: &mock.UserMock{},
	}

	routes := mux.NewRouter()
	h.InitUserRoutes(routes)

	http.ListenAndServe(":8080", routes)
}
