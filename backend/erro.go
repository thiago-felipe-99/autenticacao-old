package main

type Erro struct {
	Mensagem string
	Inicial  error
}

func (erro *Erro) Error() string {
	return erro.Mensagem
}
