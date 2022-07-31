// Package http criar uma API http da aplicação
package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thiago-felipe-99/autenticacao/backend/erro"
	"github.com/thiago-felipe-99/autenticacao/backend/logica"
)

// Usuário representa as possíveis rotas de um Usuário.
type Usuário struct {
	Lógica *logica.Usuário
}

func (usuário *Usuário) pegarID(ginC *gin.Context) {
	log.Println("Pegando ID")
}

func (usuário *Usuário) pegarBody(ginC *gin.Context) {
	log.Println("Pegando body")
}

// Criar representa a rota para criar um usuário.
func (usuário *Usuário) Criar(ginC *gin.Context) {
	ginC.JSON(http.StatusCreated, gin.H{"mensagem": "criado"})
}

// Atulizar representa a rota para atualizar um usuário.
func (usuário *Usuário) Atulizar(ginC *gin.Context) {
	ginC.JSON(http.StatusOK, gin.H{"mensagem": "atualizado"})
}

// PegarPorID representa a rota para pegar um usuário por ID.
func (usuário *Usuário) PegarPorID(ginC *gin.Context) {
	ginC.JSON(http.StatusOK, gin.H{"mensagem": "pego"})
}

// Deletar representa a rota para deletar um usuário.
func (usuário *Usuário) Deletar(ginC *gin.Context) {
	ginC.JSON(http.StatusOK, gin.H{"mensagem": "deletado"})
}

// IniciarServidor iniciar um servidor http.
func IniciarServidor(usuário *logica.Usuário) *erro.Erro {
	rotasUsuário := Usuário{
		Lógica: usuário,
	}

	servidor := gin.Default()

	servidor.POST("/usuário", rotasUsuário.pegarBody, rotasUsuário.Criar)
	servidor.GET("/usuário/:id", rotasUsuário.pegarID, rotasUsuário.PegarPorID)
	servidor.PUT(
		"/usuário/:id",
		rotasUsuário.pegarID,
		rotasUsuário.pegarBody,
		rotasUsuário.Atulizar,
	)
	servidor.DELETE("/usuário/:id", rotasUsuário.pegarID, rotasUsuário.Deletar)

	if err := servidor.Run(); err != nil {
		return &erro.Erro{
			Mensagem: "Erro ao iniciar o servidor",
			Inicial:  err,
		}
	}

	return nil
}
