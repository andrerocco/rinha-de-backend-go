package models

import (
	"github.com/google/uuid"
)

func NewPessoa(apelido string, nome string, nascimento string) Pessoa {
	return Pessoa{
		Apelido:    apelido,
		Nome:       nome,
		Nascimento: nascimento,
	}
}

type Pessoa struct {
	Id         uuid.UUID `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento string    `json:"nascimento"`
	Stack      []string  `json:"stack,omitempty"`
}
