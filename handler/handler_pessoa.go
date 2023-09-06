package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/andrerocco/rinha-de-backend-go/models"
)

func WriteJSON(w http.ResponseWriter, statusCode int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}

func HandlePessoa(w http.ResponseWriter, r *http.Request) {
	// Remove as barras extras no final do URL, se houver
	url := strings.TrimRight(r.URL.Path, "/")

	switch r.Method {
	case http.MethodGet:
		if url == "/pessoas" {
			// Entra nessa condição quando a URL é "/pessoas" ou "/pessoas/"
			// listarPessoas(w, r)
		} else if strings.HasPrefix(url, "/pessoas/") {
			// Entra nessa condição quando a URL é "/pessoas/123" ou "/pessoas/123/"
			getPessoa(w, r)
		} else {
			http.Error(w, "Rota não encontrada", http.StatusNotFound)
		}
	case http.MethodPost:
		postPessoa(w, r)
	default:
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
	}
}

func getPessoa(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("> Request received at '/pessoas/%s'\n", r.URL.Path[len("/pessoas/"):])

	pessoa := models.NewPessoa("andrerocco", "André Rocco", "1989-01-01")

	WriteJSON(w, http.StatusOK, pessoa)
	// return WriteJSON(w, http.StatusOK, pessoa)
}

func postPessoa(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST /pessoas\n")
	json.NewEncoder(w).Encode("POST /pessoas")
}

/* func postPessoa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var novaPessoa Pessoa
	// Decodifica o corpo do JSON da requisição para novaPessoa
	err := json.NewDecoder(r.Body).Decode(&novaPessoa)
	if err != nil {
		http.Error(w, "Erro ao ler dados da requisição", http.StatusBadRequest)
		return
	}

	// Valida os dados da requisição
	if novaPessoa.Apelido == "" || novaPessoa.Nome == "" || novaPessoa.Nascimento == "" {
		http.Error(w, "Campos obrigatórios não preenchidos", http.StatusUnprocessableEntity)
		return
	}

	// Verifique se o apelido já existe (simule a verificação no banco de dados)
	for _, pessoa := range pessoas {
		if pessoa.Apelido == novaPessoa.Apelido {
			http.Error(w, "O 'apelido' fornecido já existe", http.StatusConflict)
			return
		}
	}

	// Gera um novo ID
	newId, _ := uuid.NewRandom()
	novaPessoa.Id = newId

	// Adiciona a nova pessoa à lista (simule o armazenamento no banco de dados)
	pessoas = append(pessoas, novaPessoa)

	// Retorna o status code 201 (Created) e o cabeçalho "Location" com o URL da pessoa criada
	w.Header().Set("Location", fmt.Sprintf("/pessoas/%s", newId.String()))
	w.WriteHeader(http.StatusCreated)

	// Retorna um JSON com os detalhes da pessoa criada
	json.NewEncoder(w).Encode(novaPessoa)
}

func getPessoa(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	// Extrai o ID da URL
	idDaURL := r.URL.Path[len("/pessoas/"):]
	if idDaURL == "" {
		http.Error(w, "ID da pessoa não fornecido na URL", http.StatusBadRequest)
		return
	}

	// Encontre a pessoa com o ID correspondente (simule uma busca no banco de dados)
	var pessoaEncontrada Pessoa
	for _, pessoa := range pessoas {
		if pessoa.Id.String() == idDaURL {
			pessoaEncontrada = pessoa
			break
		}
	}

	// Verifique se a pessoa foi encontrada
	if pessoaEncontrada.Id.String() == "" {
		http.Error(w, "Pessoa não encontrada", http.StatusNotFound)
		return
	}

	// Retorna os detalhes da pessoa encontrada
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pessoaEncontrada)
}
*/
