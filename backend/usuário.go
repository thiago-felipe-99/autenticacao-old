package main

import "github.com/google/uuid"

// Usuário representa um usuário da aplicação
type Usuário struct {
	ID      uuid.UUID
	Nome    string
	Apelido string
	Email   string
	Senha   string
}

// Criar cria um usuário na aplicação.
func Criar(nome, senha, apelido string) (*Usuário, *Erro) {
	// verificar se existe existe(apelido, email)
	// verificar se senha é válida
	// criar usário em data
	return nil, nil
}

// BuscarPorNome todos os usuários com nome parecidos.
func BuscarPorNome(nome string) (*[]Usuário, *Erro) {
	// Pesquisar por nome em data.
	return nil, nil
}

// BuscarPorID retorna um usuário pelo seu ID.
func BuscarPorID(id string) (*Usuário, *Erro) {
	// verifcar se ID é um UUID válido
	// Pesquisar o ID em data
	return nil, nil
}

// AlterarSenha altera a senha do usuário da aplicação.
func AlterarSenha(id, senhaVelha, senhaNova string) *Erro {
	// verificar se ID é válido
	// Pegar usuário em data pelo ID
	// verificar se a senha do usuário é igual a senhaNova
	// verificar se a senhaNova é válida
	// alterar a senha do usuário em data
	return nil
}

// Deletar remove um usuário da aplicação.
func Deletar(id, senha string) *Erro {
	// verificar se ID é válido
	// Pegar usuário em data pelo ID
	// verificar se as senhas são iguais
	// apagar usuário
	return nil
}
