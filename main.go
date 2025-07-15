package main

import (
	"fmt"
	"go-net-ping/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Iniciando a aplicação GoNetPing...")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
