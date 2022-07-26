package main

import (
	"fmt"

	"github.com/ocionejr/go-api-rest-alura/models"
	"github.com/ocionejr/go-api-rest-alura/routes"
)

func main() {
	models.Personalidades = []models.Personalidade {
		{Id: 1, Nome: "Nome 1", Historia: "História 1"},
		{Id: 2, Nome: "Nome 2", Historia: "História 2"},
	}

	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

