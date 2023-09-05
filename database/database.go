package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Database struct {
	postgres *sql.DB
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "andre"
	dbname   = "rinha-de-backend-go"
)

// Inicializa a conexão com o banco de dados e cria as tabelas
func Init() (*Database, error) {
	// Conecta ao servidor PostgreSQL sem especificar um banco de dados
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Verifica se o banco de dados já existe
	if !databaseExists(db, dbname) {
		// Cria o banco de dados
		err = createDatabase(db)
		if err != nil {
			return nil, err
		}
	}

	// Conecta ao banco de dados criado
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("> Conexão com o banco de dados estabelecida")

	database := &Database{
		postgres: db,
	}
	fmt.Println("> Criando tabelas...")

	err = database.createTables()
	if err != nil {
		return nil, err
	}
	fmt.Println("> Tabelas criadas")

	return database, nil
}

// Fecha a conexão com o banco de dados
func (s *Database) Close() {
	s.postgres.Close()
}

// databaseExists verifica se um banco de dados existe no servidor PostgreSQL
func databaseExists(db *sql.DB, dbName string) bool {
	query := fmt.Sprintf("SELECT 1 FROM pg_database WHERE datname = '%s'", dbName)
	var result int
	err := db.QueryRow(query).Scan(&result)
	return err == nil
}

// CreateDatabase cria o banco de dados no servidor PostgreSQL
func createDatabase(db *sql.DB) error {
	query := fmt.Sprintf("CREATE DATABASE \"%s\"", dbname)
	_, err := db.Exec(query)
	return err
}

// CreateTables cria as tabelas no banco de dados
func (s *Database) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS pessoa (
	    id UUID PRIMARY KEY,
	    apelido TEXT NOT NULL,
	    nome TEXT NOT NULL,
	    nascimento TEXT NOT NULL,
	    stack TEXT[]
	);
	`

	_, err := s.postgres.Exec(query)
	return err
}
