package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/andrerocco/rinha-de-backend-go/database"
	"github.com/andrerocco/rinha-de-backend-go/models"
)

type PessoaRequest struct {
	Apelido    string   `json:"apelido" validate:"required,max=32"`
	Nome       string   `json:"nome" validate:"required,max=100"`
	Nascimento string   `json:"nascimento" validate:"required,datetime=2006-01-02"`
	Stack      []string `json:"stack" validate:"dive,max=32"`
}

type PessoaResponse struct {
	ID         string   `json:"id"`
	Apelido    string   `json:"apelido" validate:"required,max=32"`
	Nome       string   `json:"nome" validate:"required,max=100"`
	Nascimento string   `json:"nascimento" validate:"required,datetime=2006-01-02"`
	Stack      []string `json:"stack" validate:"dive,max=32"`
}

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

	var input PessoaRequest
	// Decodifica o corpo do JSON da requisição para pessoa
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		fmt.Printf("> Erro ao ler dados da requisição: %s\n", err.Error())
		http.Error(w, "Erro ao ler dados da requisição", http.StatusBadRequest)
		return
	}

	// Valida os dados da requisição - TODO Função validate()
	if input.Apelido == "" || input.Nome == "" || input.Nascimento == "" {
		http.Error(w, "Campos obrigatórios não preenchidos", http.StatusUnprocessableEntity)
		return
	}

	id := uuid.New()
	pessoaCreated := models.Pessoa{
		Id:         id,
		Apelido:    input.Apelido,
		Nome:       input.Nome,
		Nascimento: input.Nascimento,
		Stack:      strings.Join(input.Stack, ","),
	}

	// O campo Apelido preciso ser único e para isso, tenta-se inserir a pessoa no banco de dados
	// Se o banco já houver uma pessoa com o mesmo apelido, o banco de dados retornará um erro
	err = p.db.InserePessoa(pessoaCreated)
	if err != nil {
		fmt.Printf("> Erro ao inserir pessoa no banco de dados: %s\n", err.Error())
		http.Error(w, "O 'apelido' fornecido já existe", http.StatusConflict)
		return
	}

	// Serializa a pessoa para JSON
	response := PessoaResponse{
		ID:         pessoaCreated.Id.String(),
		Apelido:    pessoaCreated.Apelido,
		Nome:       pessoaCreated.Nome,
		Nascimento: pessoaCreated.Nascimento,
		Stack:      strings.Split(pessoaCreated.Stack, ","),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
