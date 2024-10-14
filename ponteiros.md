# Ponteiros

Ponteiro = Endereço (de memória) = Referência

Criamos a função AumentaPreco no `produto.go`, no pacote `model` à título de exemplo

```go
func AumentaPreco(produto Produto) {
    produto.Preco = produto.Preco + 3
}
```

Na `main.go`:
- comentamos o conteúdo que estava lá anteriormente
- criamos um produto
- imprimimos o preço desse produto

```go
func main() {
  // db.InitDB()
  // StartServer()

  prod := model.Produto{Nome: "Meia", Preco: 17.99}
  fmt.Println("Preço:", prod.Preco)
}
```

Output: 
```
Preço: 17.99
```

Após, chamamos a função `AumentaPreco`, passando o produto criado

```go
func main() {
  // ...
  AumentaPreco(prod)
  fmt.Println("Preço:", prod.Preco)
}
```

Ao rodar esse código, verá que o output será o seguinte:

Output: 
```
Preço: 17.99
Preço: 17.99
```

🤯🤯🤯!

Isso acontece porque quando passamos `prod` como argumento à função `AumentaPreco`, o Go cria uma **cópia** dessa estrutura Produto e lida com essa cópia dentro da função.

Então, quando é feita a alteração no preço do produto, na verdade a nova estrutura Produto (a cópia) é que está sendo alterada.

Podemos imprimir o endereço de memória dessas variáveis produto, para entendermos que estamos lidando com produtos diferentes.

Para isso, usamos `fmt.Printf`, com o placeholder `%p` para o ponteiro. O `\n` é para adicionar uma nova linha.

Para pegar o **endereço** da variável Produto, adicionamos `&` antes do nome da variável:

```go
func AumentaPreco(produto Produto) {
	fmt.Printf("AumentaPreco - Endereço de memória do prod: %p\n", &produto)
	produto.Preco = produto.Preco + 3
}
```

```go
func main() {
  // ...

	fmt.Printf("main - Endereço de memória do prod: %p\n", &prod)
}
```

Output: (vai ser algo parecido com isso)
```
Preço: 17.99
AumentaPreco - Endereço de memória do prod: 0xc0000a00a0
Preço: 17.99
main - Endereço de memória do prod: 0xc0000a0050
```

Note que o endereço é diferente! Por isso o preço do produto criado não muda - ele não foi alterado. O produto que foi alterado foi **a cópia**.

Código até o momento: https://go.dev/play/p/uhUVP68JhDz

Para que a gente faça a modificação do produto que criamos, precisamos passar para a função `AlteraPreco` o **endereço de memória** deste produto.

Para isso, precisamos alterar o parâmetro da função para que seja um **ponteiro** de Produto - fazemos isso adicionando um `*` na frente do tipo `Produto`:

```go
func AumentaPreco(produto *Produto) {
	fmt.Printf("AumentaPreco - Endereço de memória do prod: %p\n", produto) // "&" não é mais necessário antes da variável `produto`
	produto.Preco = produto.Preco + 3
}
```

Na hora de chamar a função, para passar o endereço de memória do produto, usamos o caracter `&` antes da variável: `AumentaPreco(&prod)`

Código após alterações: https://go.dev/play/p/jx9SnkYOLO5

<br>
<br>
Outro exemplo: https://go.dev/play/p/r92c2Ylyzcu

Usa-se `*` antes do ponteiro para obter o valor para o qual o ponteiro aponta. Na função `fazerAniversario` agora estamos criando uma variável chamada `novaIdade` (inteiro) e atribuindo a ela o valor para o qual o ponteiro `i` aponta. Em seguida, aumentamos em 1 a idade. 

E para atribuir um novo valor ao ponteiro, usa-se `*` na frente da variável do tipo ponteiro (nesse exemplo, `i`): `*i = novaIdade`.

Código após as alterações: https://go.dev/play/p/jrl2MfDPsza
