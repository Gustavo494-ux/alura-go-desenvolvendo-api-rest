package routes

import (
	"log"
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
