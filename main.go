package main

import (
	"fmt"

	"github.com/Gustavo494-ux/go-api-rest/routes"
)

func main() {
	fmt.Println("Iniciando o servidor com go")
	routes.HandleResquest()
}
