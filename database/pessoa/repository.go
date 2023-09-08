package pessoa

import (
	"github.com/andrerocco/rinha-de-backend-go/database"
	"github.com/andrerocco/rinha-de-backend-go/models"
)

// O arquivo repository.go contém a definição da interface Repository e a implementação do repositório de pessoa pessoaRepository.
// O repositório é responsável por interagir com o banco de dados e realizar as operações de CRUD. O repositório é utilizado
// pelo serviço (service) para realizar as operações de banco de dados necessárias para atender as solicitações HTTP.

/*
type Pessoa struct {
	Id         uuid.UUID `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento string    `json:"nascimento"`
	Stack      []string  `json:"stack,omitempty"`
}
*/

type Repository interface {
	Save(pessoa models.Pessoa) error
	Update(pessoa models.Pessoa) error
	Delete(id string) error
	GetAll() ([]models.Pessoa, error)
	GetByID(id string) (models.Pessoa, error)
}

type repository struct {
	db *database.DB // TODO - Fix
	// cache rueidis.Client
}

func NewRepository(db *database.DB) Repository {
	return &repository{
		db: db,
	}
}

func (rep *repository) Save(pessoa models.Pessoa) error {
	return nil
}

func (rep *repository) Update(pessoa models.Pessoa) error {
	return nil
}

func (rep *repository) Delete(id string) error {
	return nil
}

func (rep *repository) GetAll() ([]models.Pessoa, error) {
	return nil, nil
}

func (rep *repository) GetByID(id string) (models.Pessoa, error) {
	return models.Pessoa{}, nil
}
