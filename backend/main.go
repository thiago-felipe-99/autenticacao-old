package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/thiago-felipe-99/autenticacao/backend/data"
	"github.com/thiago-felipe-99/autenticacao/backend/http"
	"github.com/thiago-felipe-99/autenticacao/backend/logica"
)

func main() {
	bd, err := sql.Open("mysql", "dellis:@/shud")
	if err != nil {
		log.Panicf("Erro ao criar a conexão com o banco de dados: %v", err)
	}

	defer func() {
		err := bd.Close()
		if err != nil {
			log.Printf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
	}()

	data := &data.SQL{
		Conexão: bd,
	}
	lógica := &logica.Usuário{
		BD: data,
	}

	erro := http.IniciarServidor(lógica)
	if erro != nil {
		log.Panicln(erro.Error())
	}
}
