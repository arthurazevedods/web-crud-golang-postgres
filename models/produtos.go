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

	selectDeTodosOsProdutos, err := db.Query("select * from produtos order by id")
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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db, err := db.ConectaComBancoDeDados()
	if err != nil {
		panic(err.Error())
	}

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)

	defer db.Close()

}

func DeletaProduto(id string) {
	db, err := db.ConectaComBancoDeDados()

	if err != nil {
		panic(err.Error())
	}

	deletarProduto, err := db.Prepare("delete from produtos where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarProduto.Exec(id)

	defer db.Close()
}

func EditProduct(id string) Produto {
	db, err := db.ConectaComBancoDeDados()
	if err != nil {
		panic(err.Error())
	}

	productOfDB, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for productOfDB.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = productOfDB.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade

	}
	defer db.Close()
	return produtoParaAtualizar
}

func AtualizarProduto(id int, nome, descricao string, preco float64, quantidade int) {
	db, err := db.ConectaComBancoDeDados()
	if err != nil {
		panic(err.Error())
	}

	Atualiza, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	Atualiza.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()

}
