package database

import "github.com/lib/pq"

// Pessoa representa um registro na tabela pessoas
type Pessoa struct {
	Id         string   `json:"id"`
	Apelido    string   `json:"apelido"`
	Nome       string   `json:"nome"`
	Nascimento string   `json:"nascimento"`
	Stack      []string `json:"stack,omitempty"`
}

// InserePessoa insere uma nova pessoa na tabela pessoas
func (s *Database) InserePessoa(pessoa Pessoa) error {
	query := `
	INSERT INTO pessoas (id, apelido, nome, nascimento, stack)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err := s.postgres.Exec(query, pessoa.Id, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pq.Array(pessoa.Stack))
	return err
}

// AtualizaPessoa atualiza uma pessoa existente na tabela pessoas
func (s *Database) AtualizaPessoa(pessoa Pessoa) error {
	query := `
	UPDATE pessoas
	SET apelido = $2, nome = $3, nascimento = $4, stack = $5
	WHERE id = $1
	`
	_, err := s.postgres.Exec(query, pessoa.Id, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pq.Array(pessoa.Stack))
	return err
}

// ExcluiPessoa exclui uma pessoa da tabela pessoas
func (s *Database) ExcluiPessoa(id string) error {
	query := `
	DELETE FROM pessoas
	WHERE id = $1
	`
	_, err := s.postgres.Exec(query, id)
	return err
}

// ConsultaPessoas consulta todas as pessoas na tabela pessoas
func (s *Database) ConsultaPessoas() ([]Pessoa, error) {
	query := `
	SELECT id, apelido, nome, nascimento, stack
	FROM pessoas
	`
	rows, err := s.postgres.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pessoas := []Pessoa{}
	for rows.Next() {
		var pessoa Pessoa
		err = rows.Scan(&pessoa.Id, &pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, pq.Array(&pessoa.Stack))
		if err != nil {
			return nil, err
		}
		pessoas = append(pessoas, pessoa)
	}

	return pessoas, nil
}
