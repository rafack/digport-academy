# Instruções: Aula 8 banco de dados

Sugestão:

git  checkout -b aula7  (para guardar as aulas que foram dadas até aqui)

git checkout main (voltar para o branch main)

Install postgres:

https://www.postgresql.org/download/

Escolher user e senha: digport

No pgadmin criar a base de dados digport_loja

Abrir tools -> query tools
criar a tabela de produtos:

```sql
create table produtos (
id varchar primary key,
nome varchar,
preco  float8,
descricao varchar,
imagem varchar,
quantidade int
)
```
Link para ver a documentação de tipo de dados no postgress:
https://www.postgresql.org/docs/current/datatype.html

Abrir tools -> query tools
Script para inserir os produtos:
```sql
insert into produtos (id, nome, preco, descricao,imagem, quantidade) values
('f31beae8-5439-4b1f-b1f8-fbf223d4daee','Revista moda',19.90,'Revista moda','imagem1.jpg',120),
('6c9d9b24-2a73-43e6-b8a2-9d3e87566452','Revista beleza',19.90,'Revista beleza','imagem2.jpg',100)

```

https://pkg.go.dev/github.com/lib/pq#section-readme

Instalando a biblioteca uuid:  go get github.com/google/uuid  
Instalando a biblioteca postgres: go get github.com/lib/pq  
Instalando a biblioteca mux : go get github.com/gorilla/mux 

Link sobre o gorilla mux: https://github.com/gorilla/mux

Função para conectar o banco de dados no go:

```go
func ConectaBancoDados() *sql.DB {
	connStr := "user=postgres dbname=digport_loja password=digport host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

```
