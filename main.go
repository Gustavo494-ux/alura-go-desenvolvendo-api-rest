package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Gustavo494-ux/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor com go")
	routes.HandleResquest()

	log.Fatal(http.ListenAndServe(":8000", nil))
}
