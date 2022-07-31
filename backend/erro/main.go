// Package erro descreve como deve ser tratado os erros da aplicação
package erro

// Erro representa um erro da aplicação.
type Erro struct { //nolint:errname
	Mensagem string
	Inicial  error
}

// Error torna a struct Erro compatível com a interface error.
func (erro *Erro) Error() string {
	mensagem := erro.Mensagem

	if erro.Inicial != nil {
		mensagem += "\n" + erro.Inicial.Error()
	}

	return mensagem
}
