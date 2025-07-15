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

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			// Valor padrão para o host, caso o mesmo não seja fornecido
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			// Nome do comando que será utilizado na linha de comando
			Name:  "ip",
			Usage: "Busca o IP de um servidor",
			// Definindo a flag para o host. Flag é um parâmetro que pode ser passado na linha de comando
			Flags:  flags,
			Action: buscarIps,
		},
		{
			// Nome do comando que será utilizado na linha de comando
			Name:  "servers",
			Usage: "Busca o nome de servidores na internet",
			// Definindo a flag para o host. Flag é um parâmetro que pode ser passado na linha de comando
			Flags:  flags,
			Action: buscarServers,
		},
	}

	return app
}

func buscarIps(client *cli.Context) {
	host := client.String("host")

	ips, erro := net.LookupIP(host)

	verificarErro(erro)

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServers(client *cli.Context) {
	host := client.String("host")

	servers, erro := net.LookupNS(host)

	verificarErro(erro)

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func verificarErro(erro error) {
	if erro != nil {
		log.Fatal(erro)
	}
}
