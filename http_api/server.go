package http_api

import (
	"fmt"
	"net/http"

	"github.com/andrerocco/rinha-de-backend-go/database"
)

type APIServer struct {
	port string
	db   *database.DB
}

func NewAPIServer(port string, db *database.DB) *APIServer {
	return &APIServer{
		port: port,
		db:   db,
	}
}

// Inicializa o servidor na porta (especificada na instância de APIServer) e registra os handlers
// necessários para o funcionamento da API.
func (s *APIServer) Start() {
	// Initializa as rotas
	router := NewRouter(s.db)
	router.MapRoutes()

	// Inicializa o servidor
	fmt.Printf("> Servidor inicializado (:%s)\n", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", s.port), nil)
	if err != nil {
		fmt.Printf("> Erro ao inicializar o inicializar o servidor de roteamento: %s\n", err.Error())
		panic(err)
	}
}
