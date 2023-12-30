package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/database"
	"github.com/Gustavo494-ux/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var Personalidades []models.Personalidade

	database.DB.Find(&Personalidades)
	json.NewEncoder(w).Encode(Personalidades)
}

func RetornarUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var Personalidade models.Personalidade

	database.DB.First(&Personalidade, id)
	json.NewEncoder(w).Encode(Personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade

	err := json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	if err != nil {
		log.Fatal(err)
	}

	database.DB.Create(&novaPersonalidade)

	json.NewEncoder(w).Encode(novaPersonalidade)
}
