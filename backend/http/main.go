// Package http criar uma API http da aplicação
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiago-felipe-99/autenticacao/backend/logica"
)

// Usuário representa as possíveis rotas de um Usuário.
type Usuário struct {
	Lógica *logica.Usuário
}

// Criar representa a rota para criar um usário.
func (usário *Usuário) Criar(ginC *gin.Context) {
	ginC.JSON(http.StatusCreated, gin.H{"mensagem": "criado"})
}

// Atulizar representa a rota para atualizar um usário.
func (usário *Usuário) Atulizar(ginC *gin.Context) {
	ginC.JSON(http.StatusOK, gin.H{"mensagem": "atualizado"})
}

// PegarPorID representa a rota para pegar um usário por ID.
func (usário *Usuário) PegarPorID(ginC *gin.Context) {
	ginC.JSON(http.StatusOK, gin.H{"mensagem": "pego"})
}

// Deletar representa a rota para deletar um usário.
func (usário *Usuário) Deletar(ginC *gin.Context) {
	ginC.JSON(http.StatusOK, gin.H{"mensagem": "deletado"})
}
