package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	// models.Personalidade = []models.Personalidade{
	// 	{Nome: "Nome 1", Historia: "Historia 1"},
	// 	{Nome: "Nome 2", Historia: "Historia 2"},
	// }

	json.NewEncoder(w).Encode(models.Personalidades)
}
