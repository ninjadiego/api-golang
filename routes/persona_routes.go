package routes

import (
	"github/ninjadiego/demo-crud-rest-api-go/controller"

	"github.com/gorilla/mux"
)

func SetPersonaRoutes(router *mux.Router) {
	SubRoute := router.PathPrefix("/persona/api").Subrouter()
	SubRoute.HandleFunc("/all", controller.GetAll).Methods("GET")
	SubRoute.HandleFunc("/save", controller.Save).Methods("POST")
	SubRoute.HandleFunc("/delete", controller.Delete).Methods("POST")
	SubRoute.HandleFunc("/find/{id}", controller.Get).Methods("GET")

}
