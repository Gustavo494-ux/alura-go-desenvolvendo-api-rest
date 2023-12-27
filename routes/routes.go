package routes

import (
	"log"
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
