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
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)

	log.Fatal(http.ListenAndServe(":8000", r))
}
