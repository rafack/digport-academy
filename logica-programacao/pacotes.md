# Pacotes
Os programas em Go são dividos em pacotes.
Um pacote é uma coleção de arquivos **relacionados** dentro de uma _mesma pasta/diretório_. 
Essa separação serve para organização do código.

:eyes: Funções, tipos, variáveis e contantes definidos dentro de um arquivo do pacote são visíveis para todos os outros arquivos neste mesmo pacote.
Se desejamos que pacotes _diferentes_ acessem determinada propriedade ou função, usamos letra maiúscula, como explicado [aqui](funcoes#modificadores-de-acesso)

São criados com a palavra `package` na primeira linha do arquivo, em seguida de um nome.
```go
package welcome
```
[Exemplo pacote welcome](./funcoes/welcome)

## Utilizando pacotes: `import`
Os imports sempre vem após a declaração do pacote
### Importando pacotes do seu próprio projeto
O pacote `welcome` está sendo utilizado na [main](./funcoes/main.go). Para que isso seja possível, foi necessário fazer a **importação** do pacote naquele arquivo, adicionando o seguinte:
`import "<caminho-módulo>/<nome-pacote>"`

Sendo o caminho do módulo no exemplo `github.com/rafack/digport-academy/funcoes` e o nome do pacote `welcome`, temos
```go
import "github.com/rafack/digport-academy/funcoes/welcome"
```

> Sobre [Módulo](../modulo.md)

### Import único | Múltiplos imports
Quando importamos apenas um pacote no arquivo Go: `import "fmt"`. Tudo na mesma linha.

Quando importamos 2 ou mais utilizamos `( )`, dessa forma:
```go
package main

import (
  "fmt"
  "rsc.io/quote/v4"
)
```

### Pacotes padrão da linguagem Go
O Go tem diversos pacotes que são instalados juntamente com a linguagem. É a chamada [Standard Library](https://pkg.go.dev/std).

Dentre os pacotes mais utilizados da biblioteca padrão do Go está o `fmt` (mas existem muitos outros). Por ser um pacote nativo da linguagem, já instalado previamente, podemos simplesmente adicionar `import "fmt"` a um arquivo Go e utilizar suas funcionalidades.

### Importando pacotes de projetos externos
No site [pkg.go.dev](https://pkg.go.dev/) podemos pesquisar pacotes/bibliotecas para as mais diversas finalidades. Nesse lugar encontramos tanto as bibliotecas padrão/nativas do Go quanto as desenvolvidas por outras pessoas. Para este útlimo caso, precisamos realizar a instalação da biblioteca.

Ao encontrarmos uma que desejamos utilizar:
1. Copiamos o caminho do pacote

Exemplo:

![image](https://github.com/rafack/digport-academy/assets/70387077/2c3d8f9a-01a4-412f-85e6-bf3996c48f16)

2. Rodamos o seguinte comando para instalar o pacote:
```shell
go get <caminho-pkg>
```

No arquivo desejado, adicionamos o `import` e então podemos utilizar os recursos públicos (exportados) desse projeto. Para isso, usamos o mesmo caminho usado para o `go get`, entre aspas:
```go
package main

import (
  "fmt"
  "rsc.io/quote/v4"
)

func main(){
  fmt.Println(quote.Go())
}
```
Usamos o **nome do pacote** seguido de `.` para utilizar uma de suas propriedades ou funções. O nome é o que está escrito em negrito na imagem do exemplo acima.

[Mais um exemplo de uso de dependência externa](dep-externa)

## Aliases
Em qualquer momento que sentirmos necessidade, podemos colocar um "apelido" (alias) nos pacotes, de forma a dar mais sentido ao código e facilitar a leitura.
Exemplo de uso:

```go
package main

import (
  "fmt"
  geradorFrases "rsc.io/quote/v4"
)

func main(){
  fmt.Println(geradorFrases.Go())
}
```
