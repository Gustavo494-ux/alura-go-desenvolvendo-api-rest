package routes

import (
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/controllers"
)

func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
}
