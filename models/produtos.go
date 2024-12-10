package models

import (
	"fmt"

	"github.com/arthurazevedods/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db, error := db.ConectaComBancoDeDados()
	if error != nil {
		fmt.Print("error:", error)
	}

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(fmt.Sprintf("Erro ao consultar os produtos: %v", err))
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(fmt.Sprintf("Erro ao ler dados dos produtos: %v", err))
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}
