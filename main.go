package main

import (
	"fmt"
	"net/http"

	"github.com/andrerocco/rinha-de-backend-go/database"
	"github.com/andrerocco/rinha-de-backend-go/handler"
)

type APIServer struct {
	port int
}

// Cria um novo objeto APIServer e retorna o ponteiro
func NewAPIServer(port int) *APIServer {
	return &APIServer{port: port}
}

// Inicializa o servidor na porta (especificada no objeto APIServer) e registra os handlers
func (s *APIServer) Start() {
	// Registra os handlers
	http.HandleFunc("/pessoas/", handler.HandlePessoa)

	// Inicializa o servidor
	fmt.Printf("Servidor inicializado (:%d)\n", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	_, err := database.Init()
	if err != nil {
		panic(err)
	}

	server := NewAPIServer(8080)
	server.Start()
}
