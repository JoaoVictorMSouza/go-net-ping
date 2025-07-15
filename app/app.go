package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar irá a aplicação de linha de comando pronta para ser executada.
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "GoNetPing"
	app.Usage = "Busca IP's e nomes de servidores na internet"

	app.Commands = []cli.Command{
		{
			// Nome do comando que será utilizado na linha de comando
			Name:  "ip",
			Usage: "Busca o IP de um servidor",
			// Definindo a flag para o host. Flag é um parâmetro que pode ser passado na linha de comando
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					// Valor padrão para o host, caso o mesmo não seja fornecido
					Value: "google.com",
				},
			},
			Action: buscarIps,
		},
	}

	return app
}

func buscarIps(client *cli.Context) {
	host := client.String("host")

	ips, erro := net.LookupIP(host)

	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
