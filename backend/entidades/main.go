package entidades

import "github.com/google/uuid"

// Usuário representa um usuário da aplicação.
type Usuário struct {
	ID      uuid.UUID
	Nome    string
	Apelido string
	Email   string
	Senha   string
}
