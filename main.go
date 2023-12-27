package main

import (
	"fmt"

	"github.com/Gustavo494-ux/go-api-rest/models"
	"github.com/Gustavo494-ux/go-api-rest/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}

	fmt.Println("Iniciando o servidor com go")
	routes.HandleResquest()
}
