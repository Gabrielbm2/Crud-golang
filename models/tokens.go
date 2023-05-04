package models

import (
	"gocrud/db"
)

type Token struct {
	Id    int
	Nome  string
	Preco float64
}

func BuscaTodosOsTokens() []Token {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsTokens, err := db.Query("select * from tokens")
	if err != nil {
		panic(err.Error())
	}

	p := Token{}
	tokens := []Token{}

	for selectDeTodosOsTokens.Next() {
		var id int
		var nome string
		var preco float64

		err = selectDeTodosOsTokens.Scan(&id, &nome, &preco)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Preco = preco

		tokens = append(tokens, p)
	}
	defer db.Close()
	return tokens
}
func CriaNovoToken(nome string, preco float64) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into tokens(nome, preco) values($1, $2)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, preco)
	defer db.Close()

}

func DeletaToken(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOToken, err := db.Prepare("delete from tokens where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOToken.Exec(id)
	defer db.Close()

}

func EditaToken(id string) Token {
	db := db.ConectaComBancoDeDados()

	tokenDoBanco, err := db.Query("select * from tokens where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	tokenParaAtualizar := Token{}

	for tokenDoBanco.Next() {
		var id int
		var nome string
		var preco float64

		err = tokenDoBanco.Scan(&id, &nome, &preco)
		if err != nil {
			panic(err.Error())
		}
		tokenParaAtualizar.Id = id
		tokenParaAtualizar.Nome = nome
		tokenParaAtualizar.Preco = preco
	}
	defer db.Close()
	return tokenParaAtualizar
}

func AtualizaToken(id int, nome string, preco float64) {
	db := db.ConectaComBancoDeDados()

	AtualizaToken, err := db.Prepare("update tokens set nome=$1, preco=$2, where id=$3")
	if err != nil {
		panic(err.Error())
	}
	AtualizaToken.Exec(nome, preco, id)
	defer db.Close()
}
