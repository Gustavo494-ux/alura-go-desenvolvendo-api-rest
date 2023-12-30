package routes

import (
	"log"
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.CriaUmaNovaPersonalidade).Methods("POST")
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornarUmaPersonalidade).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletarPersonalidade).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
