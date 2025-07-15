# Go NetPing

**Go NetPing** é uma ferramenta de linha de comando simples, desenvolvida em **Golang**, que permite realizar duas operações de rede essenciais:

- **Buscar o endereço IP** de um domínio ou servidor web.
- **Obter o nome do servidor** com base na hospedagem do endereço fornecido.

Esse projeto marca o meu primeiro contato com a linguagem Go e serve como um aprendizado prático sobre a construção de aplicações de rede simples.

## Funcionalidades

- **Buscar IP**: Obtém o endereço IP de um domínio ou servidor web.
- **Resolver Nome de Servidor**: Retorna o nome do servidor com base no endereço IP fornecido.

> O endereço **google.com** é o valor **default** para os comandos, ou seja, caso não seja informado nenhum parâmetro, **google.com** será utilizado.

## Bibliotecas Utilizadas

- [**urfave/cli**](https://github.com/urfave/cli): Usada para criar a interface de linha de comando.
- [**net**](https://pkg.go.dev/net): Biblioteca padrão do Go para operações de rede.
- [**log**](https://pkg.go.dev/log): Usada para gerenciamento de logs e mensagens de erro.

## Como Usar

1. Clone o repositório:

   ```bash
   git clone https://github.com/seu-usuario/go-net-ping.git
   cd go-net-ping

2. Instale as dependências:

   ```bash
   go mod tidy

3. Execute a aplicação via linha de comando:

   ### Buscar IP
   - Para buscar o IP de um domínio (caso o parâmetro não seja passado, google.com será usado por padrão):
     ```bash
     go run main.go ip --host google.com
     ```  
      Exemplo sem o parâmetro, usando o valor padrão:
     ```bash
     go run main.go ip
     ```
     Exemplo de saída
     ```bash
      142.250.72.14
     ```

   ### Resolver Nome do Servidor
   - Para resolver o nome do servidor baseado no IP fornecido (novamente, google.com será usado por padrão se o parâmetro não for especificado):
     ```bash
     go run main.go servers --host google.com
     ```  
      Exemplo sem o parâmetro, usando o valor padrão:
     ```bash
     go run main.go servers
     ```
     Exemplo de saída
     ```bash
     lga34s16-in-f14.1e100.net
     ```


 ## Contribuições
Este é um projeto pessoal para aprendizado, mas contribuições são sempre bem-vindas!
