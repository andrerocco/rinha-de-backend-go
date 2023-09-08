package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/andrerocco/rinha-de-backend-go/models"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Storage interface {
	InserePessoa(pessoa models.Pessoa) error
	AtualizaPessoa(pessoa models.Pessoa) error
	ExcluiPessoa(id string) error
	ConsultaPessoas() ([]models.Pessoa, error)
}

type DB struct {
	postgres *sql.DB
}

// Inicializa a conexão com o banco de dados e cria as tabelas
func Connect() (*DB, error) {
	// Conecta ao banco de dados criado
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	postgres, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println("> Conexão com o banco de dados aberta")

	err = postgres.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("> Conexão com o banco de dados estabelecida com sucesso")

	db := &DB{
		postgres: postgres,
	}

	fmt.Println("> Criando tabelas...")
	err = db.createTables()
	if err != nil {
		return nil, err
	}
	fmt.Println("> Tabelas criadas")

	return db, nil
}

// Fecha a conexão com o banco de dados
func (s *DB) Close() {
	s.postgres.Close()
}

// CreateTables cria as tabelas no banco de dados
func (s *DB) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS pessoa (
	    id UUID PRIMARY KEY,
	    apelido TEXT UNIQUE NOT NULL,
	    nome TEXT NOT NULL,
	    nascimento TEXT NOT NULL,
	    stack TEXT[]
	);
	`

	_, err := s.postgres.Exec(query)
	return err
}

func (s *DB) InserePessoa(pessoa models.Pessoa) error {
	// Converte o slice de strings para um array de strings
	stack := pq.Array(pessoa.Stack)

	_, err := s.postgres.Exec(
		"INSERT INTO pessoa (id, apelido, nome, nascimento, stack) VALUES ($1, $2, $3, $4, $5)",
		pessoa.Id, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, stack,
	)
	return err
}

func (s *DB) AtualizaPessoa(pessoa models.Pessoa) error {
	query := `
		UPDATE pessoa
		SET apelido = $1, nome = $2, nascimento = $3, stack = $4
		WHERE id = $5
	`

	_, err := s.postgres.Exec(query, pessoa.Apelido, pessoa.Nome, pessoa.Nascimento, pessoa.Stack, pessoa.Id)
	return err
}

func (s *DB) ExcluiPessoa(id string) error {
	query := `
		DELETE FROM pessoa
		WHERE id = $1
	`

	_, err := s.postgres.Exec(query, id)
	return err
}

func (s *DB) ConsultaPessoas() ([]models.Pessoa, error) {
	query := `
		SELECT id, apelido, nome, nascimento, stack
		FROM pessoa
	`

	rows, err := s.postgres.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pessoas []models.Pessoa
	for rows.Next() {
		var pessoa models.Pessoa
		err := rows.Scan(&pessoa.Id, &pessoa.Apelido, &pessoa.Nome, &pessoa.Nascimento, &pessoa.Stack)
		if err != nil {
			return nil, err
		}

		pessoas = append(pessoas, pessoa)
	}

	return pessoas, nil
}
