# Tema da aula 6: Busca de Produtos por Nome

No mesmo arquivo onde está a função para criar a lista de produtos, crie uma nova função para *buscar um produto por nome**.

Assim como a função para criar a lista de produtos, essa também retorna uma lista de produtos.

Essa função deve receber por parâmetro um nome.

💡 **Dica:** a função de criar a lista de produtos pode ser usada dentro dessa nova função

Recomendo voltar no conteúdo das aulas anteriores para ter ideias de como implementar!

Para testar essa função nova, pode usar esse código no `main.go`:
```go
package main

import "fmt"

func main() {
	nome := "Nome do Produto" // substituir por um nome de produto

	produtosFiltrados := produtosPorNome(nome)

	fmt.Println(produtosFiltrados)
}
```

## Resultados esperados
Exemplo de catálogo:
```go
catalogo := []model.Produto{
	{
		ID:                  "6",
		Nome:                "A menina que roubava livros",
		Preco:               39.90,
		Descricao:           "A menina que roubava livros",
		Imagem:              "livro.jpg",
		QuantidadeEmEstoque: 9,
	}
}
```
- Se informei na variável `nome` um nome de produto que EXISTE no meu catálogo (`nome := "A menina que roubava livros"`), devo ver `[{6  39.9 A menina que roubava livros livro.jpg 9}]` no terminal
- Se informei um nome que NÃO EXISTE, devo ver `[]` (array vazio)

## Extra
Quem estiver se sentindo confiante, pode tentar adicionar um query parameter `nome` no endpoint `/produtos` que criamos na aula de hoje.

O query parameter vai dentro da função handler que criamos (`produtosHandler`).

Precisamos usar o parâmetro recebido na request para chamar a função criada

Referência: https://pkg.go.dev/net/url#URL.Query

### Para testar:
- Rodar `go run .`
- Fazer a requisição GET no browser:
	- acessar no navegador `localhost:8080/produtos?nome=Nome do Produto`

❗ **Importante**: ao acessar `/produtos` - sem informar o parâmetro `nome` - o comportamento deve continuar o mesmo!, ou seja, retornar o catálogo todo.
