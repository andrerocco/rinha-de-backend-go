package http_api

import (
	"net/http"
	"strings"

	"github.com/andrerocco/rinha-de-backend-go/database"
	handler "github.com/andrerocco/rinha-de-backend-go/http_api/handlers"
)

type router struct {
	db *database.DB
	// eng   *gin.Engine
	// rg    *gin.RouterGroup
	// db    *pgxpool.Pool
	// cache rueidis.Client
}

// Cria um novo objeto router e retorna o ponteiro
func NewRouter(db *database.DB) *router {
	return &router{
		db: db,
	}
}

func (rtr *router) buildPessoaRoutes(w http.ResponseWriter, req *http.Request) {
	handler := handler.NewPessoaControler(rtr.db)

	// Remove as barras extras no final do URL, se houver
	url := strings.TrimRight(req.URL.Path, "/")

	switch req.Method {
	case http.MethodGet:
		if url == "/pessoas" {
			// Entra nessa condição quando a URL é "/pessoas" ou "/pessoas/"
			handler.GetAllPessoas(w, req)
		} else if strings.HasPrefix(url, "/pessoas/") {
			// Entra nessa condição quando a URL é "/pessoas/123" ou "/pessoas/123/"
			// getPessoa(w, req)
		} else {
			http.Error(w, "Rota não encontrada", http.StatusNotFound)
		}
	case http.MethodPost:
		handler.PostPessoa(w, req)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func (rtr *router) MapRoutes() {
	http.HandleFunc("/pessoas", rtr.buildPessoaRoutes)
	http.HandleFunc("/pessoas/", rtr.buildPessoaRoutes)
}
