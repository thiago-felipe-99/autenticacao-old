// Package logica representa quais são as lógicas da aplicação.
package logica

import (
	"github.com/thiago-felipe-99/autenticacao/backend/data"
	"github.com/thiago-felipe-99/autenticacao/backend/entidades"
	"github.com/thiago-felipe-99/autenticacao/backend/erro"
)

// Usuário representa as possíveis lógicas que a entidade Usuário
// pode ter.
type Usuário struct {
	BD *data.SQL
}

// Criar cria um usuário na aplicação.
func (lógica *Usuário) Criar(nome, senha, apelido string) (*entidades.Usuário, *erro.Erro) {
	// verificar se existe existe(apelido, email)
	// verificar se senha é válida
	// criar usário em data
	return nil, nil
}

// BuscarPorNome todos os usuários com nome parecidos.
func (lógica *Usuário) BuscarPorNome(nome string) (*[]entidades.Usuário, *erro.Erro) {
	// Pesquisar por nome em data.
	return nil, nil
}

// BuscarPorID retorna um usuário pelo seu ID.
func (lógica *Usuário) BuscarPorID(id string) (*entidades.Usuário, *erro.Erro) {
	// verifcar se ID é um UUID válido
	// Pesquisar o ID em data
	return nil, nil
}

// AlterarSenha altera a senha do usuário da aplicação.
func (lógica *Usuário) AlterarSenha(id, senhaVelha, senhaNova string) *erro.Erro {
	// verificar se ID é válido
	// Pegar usuário em data pelo ID
	// verificar se a senha do usuário é igual a senhaNova
	// verificar se a senhaNova é válida
	// alterar a senha do usuário em data
	return nil
}

// Deletar remove um usuário da aplicação.
func (lógica *Usuário) Deletar(id, senha string) *erro.Erro {
	// verificar se ID é válido
	// Pegar usuário em data pelo ID
	// verificar se as senhas são iguais
	// apagar usuário
	return nil
}
