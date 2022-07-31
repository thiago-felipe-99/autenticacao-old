// Package data representa quais são as interações da aplicação com os
// sistemas de amarzenamento.
package data

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/thiago-felipe-99/autenticacao/backend/entidades"
	"github.com/thiago-felipe-99/autenticacao/backend/erro"
)

// SQL represnta uma conexão com um banco de dados do tipo SQL.
type SQL struct {
	Conexão *sql.DB
}

// CriarUsuário adionca um usuário no banco de dados.
func (bd *SQL) CriarUsuário(ctx context.Context, usuário entidades.Usuário) *erro.Erro {
	_, err := bd.Conexão.ExecContext(
		ctx,
		"INSER INTO usuário (id, nome, apelido, email, senha) VALUES ($1, $2, $3, $4, $5)",
		usuário.ID,
		usuário.Nome,
		usuário.Apelido,
		usuário.Email,
		usuário.Senha,
	)
	if err != nil {
		return &erro.Erro{
			Mensagem: "Erro ao adicionar usuário no Banco De Dados",
			Inicial:  err,
		}
	}

	return nil
}

// AtualizarUsuário atualiza um usuário no banco de Dados.
func (bd *SQL) AtualizarUsuário(
	ctx context.Context,
	id uuid.UUID,
	usuário *entidades.Usuário,
) *erro.Erro {
	_, err := bd.Conexão.ExecContext(
		ctx,
		"UPDATE usuário SET nome = $1, apelido = $2, email = $3, senha = $4 WHERE id = $5",
		usuário.Nome,
		usuário.Apelido,
		usuário.Email,
		usuário.Senha,
		id,
	)
	if err != nil {
		return &erro.Erro{
			Mensagem: "Erro ao atualizar usuário no Banco De Dados",
			Inicial:  err,
		}
	}

	return nil
}

// PegarUsuárioPorID retorna um usuário pelo seu ID.
func (bd *SQL) PegarUsuárioPorID(ctx context.Context, id uuid.UUID) (
	*entidades.Usuário,
	*erro.Erro,
) {
	var usuário *entidades.Usuário

	query := bd.Conexão.QueryRowContext(
		ctx,
		"SELECT (nome, apelido, email, senha) WHERE id = $1",
		id,
	)

	err := query.Scan(usuário.Nome, usuário.Apelido, usuário.Email, usuário.Senha)
	if err != nil {
		return nil, &erro.Erro{
			Mensagem: "Erro ao pegar o usuário por ID",
			Inicial:  err,
		}
	}

	usuário.ID = id

	return usuário, nil
}

// DeletarUsuário remove um usuário do banco de dados.
func (bd *SQL) DeletarUsuário(ctx context.Context, id uuid.UUID) *erro.Erro {
	_, err := bd.Conexão.ExecContext(ctx, "DELETE FROM usuário WHERE id = $1", id)
	if err != nil {
		return &erro.Erro{
			Mensagem: "Erro ao deletar usuário do banco de dados",
			Inicial:  err,
		}
	}

	return nil
}
