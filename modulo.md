# Módulo
Serve para identificar o projeto e suas dependências. 
Geralmente, um projeto contém apenas um módulo, que está localizado na pasta raiz (principal) do repositório - a primeira pasta que vemos quando o acessamos.

Criamos um módulo com o seguinte comando:
```shell
go mod init <caminho-do-modulo>
```
**`<caminho-do-modulo>`**: o caminho informado para o `go mod init` normalmente é o endereço 🔗 do repositório onde seu código será armazenado (sem o `https://` na frente).

A partir desse comando, um arquivo chamado `go.mod` é criado no diretório atual.

Exemplo de [`go.mod`](logica-programacao/funcoes/go.mod)

Na primeira linha do arquivo, temos a identificação do módulo - `module <caminho-do-modulo>` - e a seguir (após uma linha em branco), `go <versao>`, que diz que o projeto precisa dessa versão do Go (ou superior) para funcionar corretamente.
A versão é "preenchida" automaticamente com a versão do Go que está instalada no seu PC.

## Gerenciamento de Dependências
O módulo Go, através do arquivo `go.mod`, declara quais são as dependências do projeto (outros módulos, de outros projetos/repos) e qual a versão desses módulos o projeto utilza.
Para isso, usa a palavra `require`.

Exemplo:
Quero verificar se um CPF é válido. Para isso, posso utilizar um pacote que já faz isso, pesquisando em [pkg.go.dev](https://pkg.go.dev/).
Encontrei esse pacote `github.com/Nhanderu/brdoc` e vi que ele tem o que preciso.
➡️ No diretório do meu projeto (onde está localizado o `go.mod`), rodo o comando `go get github.com/Nhanderu/brdoc` para instalar esse pacote e poder utilizá-lo

Isso faz com que esse pacote se torne uma **dependência** do projeto. Então, quando olhamos para o `go.mod`, vemos que agora ele contém `require github.com/Nhanderu/brdoc v1.1.2`, deixando explícita essa dependência, bem como a versão do pacote "brdoc" estamos utilizando no projeto (quando não especificada, `go get` instala a mais recente).

A importância disso é garantir que se tenha exatamente as versões corretas das dependências, evitando que versões incompatíveis de bibliotecas causem conflitos e erros (uma parte do código/uma função pode se comportar de uma certa maneira na versão 1.0.0 do pacote mas na versão 2.0.0 ter um comportamento diferente, que pode causar problemas no nosso projeto).
