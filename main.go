package main

import (
	"github/ninjadiego/demo-crud-rest-api-go/commons"
	"github/ninjadiego/demo-crud-rest-api-go/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	commons.Migrate()

	router := mux.NewRouter()
	routes.SetPersonaRoutes(router)

	server := http.Server{
		Addr:    ":9000",
		Handler: router,
	}

	log.Println("servidor ejecutandose sobre el puerto 9000")
	log.Println(server.ListenAndServeTLS())

}
