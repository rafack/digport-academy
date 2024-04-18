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
	nome := "Nome do Produto"

	produtosFiltrados := produtosPorNome(nome)

	fmt.Println(produtosFiltrados)
}
```
