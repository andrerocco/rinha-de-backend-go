package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andrerocco/rinha-de-backend-go/database"
	"github.com/andrerocco/rinha-de-backend-go/models"
)

/* type PessoaRequest struct {
	Apelido    string `json:"apelido"`
	Nome       string `json:"nome"`
	Nascimento string `json:"nascimento"`
	Stack      string `json:"stack,omitempty"`
}

type PessoaResponse models.Pessoa */

type PessoaControler struct {
	db *database.DB
}

func NewPessoaControler(db *database.DB) *PessoaControler {
	return &PessoaControler{
		db: db,
	}
}

func (p *PessoaControler) GetAllPessoas(w http.ResponseWriter, req *http.Request) {
	fmt.Println("> Requisição de GET recebida em /pessoas")

	// Consulta as pessoas no banco de dados
	pessoas, err := p.db.ConsultaPessoas()
	if err != nil {
		fmt.Printf("> Erro ao consultar pessoas no banco de dados: %s\n", err.Error())
		http.Error(w, "Erro ao consultar pessoas no banco de dados", http.StatusInternalServerError)
		return
	}

	// Serializa as pessoas para JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoas)
}

func (p *PessoaControler) PostPessoa(w http.ResponseWriter, req *http.Request) {
	fmt.Println("> Requisição de POST recebida em /pessoas")

	var pessoa models.Pessoa
	// Decodifica o corpo do JSON da requisição para pessoa
	err := json.NewDecoder(req.Body).Decode(&pessoa)
	if err != nil {
		http.Error(w, "Erro ao ler dados da requisição", http.StatusBadRequest)
		return
	}

	// TODO - Adicionar UUID em pessoa

	// Valida os dados da requisição
	if pessoa.Apelido == "" || pessoa.Nome == "" || pessoa.Nascimento == "" {
		http.Error(w, "Campos obrigatórios não preenchidos", http.StatusUnprocessableEntity)
		return
	}

	// O campo Apelido preciso ser único e para isso, tenta-se inserir a pessoa no banco de dados
	// Se o banco já houver uma pessoa com o mesmo apelido, o banco de dados retornará um erro
	err = p.db.InserePessoa(pessoa)
	if err != nil {
		fmt.Printf("> Erro ao inserir pessoa no banco de dados: %s\n", err.Error())
		http.Error(w, "O 'apelido' fornecido já existe", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pessoa)
}
