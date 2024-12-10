# Projeto CRUD em Golang com PostgreSQL

Este é um projeto simples de CRUD (Create, Read, Update, Delete) desenvolvido em Golang, utilizando PostgreSQL como banco de dados. O objetivo deste projeto é demonstrar como criar uma API RESTful básica em Golang com operações CRUD para gerenciar recursos em um banco de dados PostgreSQL.

## Tecnologias Utilizadas

- **Golang**: Linguagem de programação para backend.
- **PostgreSQL**: Banco de dados relacional utilizado para armazenar os dados.
- **Godotenv**: Biblioteca para carregar variáveis de ambiente de um arquivo `.env`.
- **pq**: Driver PostgreSQL para Golang.

## Estrutura do Projeto

```bash

├── main.go            # Arquivo principal do servidor
├── controllers/       # 
│   └── produtos.go    # 
├── db/                # Pacote para configuração do banco de dados
│   └── db.go          # Funções de conexão e manipulação do banco de dados
├── models/            # Pacote com as estruturas de dados (modelos)
│   └── produtos.go    # Modelo de dados dos produtos
├── routes/            # 
│   └── routes.go      # 
├── templates/         # Front-end
│   └── index.html     # 
│   └── new.html       #
├── .env               # Arquivo com as variáveis de ambiente
├── go.mod             # Dependências do projeto
└── README.md          # Este arquivo
```

## Instalação
1. Clone este repositório para sua máquina local:
```bash
    git clone https://github.com/arthurazevedods/web-crud-golang-postgres.git
    cd projeto-crud-golang-postgres
```
2. Crie um arquivo .env na raiz do projeto com as seguintes variáveis:
```bash
    DATABASE_HOST=localhost
    DATABASE_PORT=5432
    DATABASE_USERNAME=seu_usuario
    DATABASE_PASSWORD=sua_senha
    DATABASE_DBNAME=seu_banco
```
3. Instale as dependências do projeto:
```bash
    go mod tidy
```
4. Crie seu banco de dados no PostgreSQL e adicione os produtos
5. Execute o projeto:
```bash
    go run main.go
```

## Contribuindo
- Faça um fork deste repositório.
- Crie uma branch para suas mudanças (git checkout -b feature/nova-funcionalidade).
- Faça o commit das suas mudanças (git commit -am 'Adiciona nova funcionalidade').
- Envie para o branch principal (git push origin feature/nova-funcionalidade).
- Abra um Pull Request.
