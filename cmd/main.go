package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andrerocco/rinha-de-backend-go/handler"
)

type APIServer struct {
	port string
}

// Cria um novo objeto APIServer e retorna o ponteiro
func NewAPIServer(port string) *APIServer {
	return &APIServer{port: port}
}

// Inicializa o servidor na porta (especificada no objeto APIServer) e registra os handlers
func (s *APIServer) Start() {
	// Registra os handlers
	http.HandleFunc("/pessoas/", handler.HandlePessoa)

	// Inicializa o servidor
	fmt.Printf("> Servidor inicializado (:%s)\n", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
	if err != nil {
		fmt.Printf("> Erro ao inicializar o inicializar o servidor de roteamento: %s\n", err.Error())
		panic(err)
	}
}

func main() {
	/* _, err := database.Init()
	if err != nil {
		panic(err)
	} */

	// Get docker env vars
	server := NewAPIServer(os.Getenv("API_PORT"))
	server.Start()
}
