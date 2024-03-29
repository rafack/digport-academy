# Pacotes
Os programas em Go são dividos em pacotes.
Um pacote é uma coleção de arquivos relacionados dentro de uma mesma pasta/diretório. 
Essa separação serve para organização do código.

:eyes: Funções, tipos, variáveis e contantes definidos dentro de um arquivo do pacote são visíveis para todos os outros arquivos neste mesmo pacote.

São criados com a palavra `package` na primeira linha do arquivo, em seguida de um nome.
```go
package welcome
```
[Exemplo pacote welcome](./funcoes/welcome). Está sendo utilizado na [main](./funcoes/main.go).
