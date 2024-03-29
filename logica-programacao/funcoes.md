# Funções
Para escrever uma função em Go, utilizamos a palavra reservada `func`. Em seguida, damos a ela um nome:
```go
func main() {

}
```
`()` são obrigatórios, é onde vão os parâmetros da função. Após, usa-se `{` para "iniciar" a função e `}` para encerrá-la.

Diversas utilidades na programação, dentre elas:
- organização do código
- melhora da legibilidade
- facilitação de testes
- atribuição de um nome significativo a um pedaço de código
> :bulb: É essencial dar **nomes descritivos** às funções, para que fique clara a intenção do código para quem o lê

## Modificadores de Acesso
> "Pre-req" para entender essa parte: [pacotes](pacotes.md)
No Go, as funções podem ser privadas ou públicas. 
Em diversas linguagens, definimos esses modificadores de acesso através de palavras reservadas como `public`, `private` e `protected`. 
Porém, o Golang faz isso de uma forma sutil, através da primeira letra do nome da função:

- se é minúscula, a função é "protegida"
```go
package welcome

import "fmt"

func sayHello(){
  fmt.Println("Hello")
}
```
A função `sayHello` não é acessível em outros pacotes

- se é maiúscula, a função é pública
```go
package welcome

import "fmt"

func SayGoodbye(){
  fmt.Println("Goodbye")
}
```
A função `SayGoodbye` pode ser acessada em pacotes diferentes do seu (pacote `welcome`):
```go
package main

func main(){
  welcome.SayGoodbye()
}
```
Estamos a invocando/chamando dentro do pacote `main`.
