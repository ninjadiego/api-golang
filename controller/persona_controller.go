package controller

import (
	"encoding/json"
	"github/ninjadiego/demo-crud-rest-api-go/commons"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetALL(writer http.ResponseWriter, request *http.Request) {
	personas := []models.persona{}
	db := commons.Getconnection()
	defer db.close()

	db.Find(&personas)
	json, _ := json.Marshal(personas)
	commons.SendResponsde(writer, http.StatusOK, json)
}

func Get(writer http.ResponseWriter, request *http.Request) {
	persona := models.persona{}

	id := mux.Vars(request)["id"]

	db := commons.Getconnection()
	defer db.close()

	db.Find(&persona, id)

	if persona.ID > 0 {
		json, _ := json.Marshal(persona)
		commons.SendResponsde(writer, http.StatusOK, json)
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}

}

func Save(writer http.ResponseWriter, request *http.Request) {
	persona := models.persona{}

	db := commons.Getconnection()
	defer db.close()

	error := json.NewDecoder(request.Body).Decode(persona)

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusBadRequest)
		return
	}

	error = db.Create(&persona).Error

	if error != nil {
		log.Fatal(error)
		commons.SendError(writer, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(persona)

	commons.SendResponsde(writer, http.StatusCreated, json)
}

func Delete(writer http.ResponseWriter, request *http.Request) {
	persona := models.persona{}

	db := commons.Getconnection()
	defer db.Close()

	id := mux.Vars(request)["id"]

	db.Find(&persona, id)

	if persona > 0 {
		db.Delete(persona)
		commons.SendResponsde(writer, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(writer, http.StatusNotFound)
	}
}
