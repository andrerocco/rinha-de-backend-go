package models

import (
	"github.com/google/uuid"
)

func NewPessoa(apelido, nome, nascimento string, stack ...string) Pessoa {
	if len(stack) == 0 {
		stack = nil
	}

	return Pessoa{
		Id:         uuid.New(),
		Apelido:    apelido,
		Nome:       nome,
		Nascimento: nascimento,
		Stack:      stack,
	}
}

type Pessoa struct {
	Id         uuid.UUID `json:"id"`
	Apelido    string    `json:"apelido"`
	Nome       string    `json:"nome"`
	Nascimento string    `json:"nascimento"`
	Stack      []string  `json:"stack,omitempty"` // TODO - Decidir se manter omitempty
}
