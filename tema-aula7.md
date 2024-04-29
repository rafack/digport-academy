# Tema da aula 7: Tratamento de Erros

Sendo o ID um identificador único dos produtos da nossa loja, não podemos permitir a criação de um novo produto com o ID igual ao de um produto já cadastrado.

> Exemplo:
>
> Request: GET /produtos
>
> Response: Lista de Produtos da minha loja
> ```json
> [
>     {
>         "id": "0",
>         "nome": "Calça legging"
>     },
>     {
>         "id": "1",
>         "nome": "Bolsa"
>     }
> ]
> ```
>
> Ao tentar cadastrar um novo produto de ID 0:
> 
> Request: POST /produtos | Body: `{"id": "0", "nome": "Boné"}`
> 
> Devo obter uma resposta assim:
> 
> Response: 400 - Bad Request | Body: `{"errorMessage": "o id é inválido"}`

Por isso, agora a função que adiciona um novo produto retorna um erro (`func adicionaProduto(produtoNovo model.Produto) error`).

## Exercício
1. Vá até função handler - onde recebemos a requisição e chamamos a função para adicionar o novo produto
2. Atribua o valor retornado por essa função (do tipo `error`) a uma variável. Por padrão, no Go o nome dessa variável é `err`
3. Verifique se houve um erro
    - Note qual é o retorno da função `adicionaProduto` caso não haja erro e o produto seja criado com sucesso. Você pode testar isso no `main.go` para entender melhor, chamando a função `adicionaProduto` e imprimindo o retorno
4. Se houver erro, use o mesmo código que escreve o status code 201 (Created) para escrever o status code 400 (Bad Request)
5. Use um código similar ao que escreve a nossa lista de produtos no corpo da resposta para, em vez disso, escrever o erro
    - No pacote `model`, em um novo arquivo, crie uma estrutura de Erro com um campo para a mensagem de erro
    - Em vez de escrever `produtos` como resposta, use a `struct` definida para criar um erro. Coloque `err.Error()` no campo da mensagem de erro. Isso vai pegar a mensagem de erro que a função `adicionaProduto` retornou
6. Teste!

Se não houve erro, o status code segue sendo 201 (Created). **O "caminho feliz" 😃 deve continuar funcionando corretamente**

> 💡 Possivelmente o vídeo indicado abaixo para estudo vá ajudar nas dúvidas durante o exercício

## Estudo

Assistam esse [vídeo sobre Funções](https://youtu.be/PPOBe49M8V0?si=8KQpsSIYN4cuGqcz) para reforçar coisas como:
- Parâmetros
- Retorno de funções

💻  **Repliquem os códigos mostrados no vídeo** para fixarem melhor e para fazerem testes (mudem os valores, os retornos, façam o código quebrar, etc)! Essa é uma parte bem importante do estudo/aprendizado

Vocês podem fazer isso no repositório que usamos pra aulas iniciais

OU

Podem usar o [Go Playground](https://go.dev/play/), como é feito no vídeo.

## Extra

Como muito bem pontuado por algumas alunas na aula, normalmente o ID é definido pelo próprio sistema e não por quem o "chama"/pelo cliente.

No momento, a nossa função pega todo o produto novo que foi passado no body da nossa requisição POST e adiciona ele, exatamente como veio, à nossa lista de produtos. Ou seja, quem está chamando nosso servidor é que está informando o valor do ID.

Para mudar isso, temos que definir um novo ID para o produto novo (no código). Então, antes de adicionar esse produto na lista (com a função `append`), devemos sobrescrever o ID da variável `produtoNovo`.

Duas possíveis abordagens:

### Usando IDs sequenciais
Sugestão de uma colega: usar o tamanho da nossa lista de produtos para definir o ID do novo produto

Supondo que iniciamos a nossa lista de produtos vazia. Quando estivermos adicionando o primeiro produto, ao verificarmos o tamanho da lista veremos que é `0`. 
Podemos usar esse valor para definir o ID do primeiro produto. Então, quando estivermos adicionando o segundo produto, o ID deste vai ser `1` (tamanho da lista após a inserção do primeiro produto), e assim por diante.

A função usada para obter o tamanho da lista é `len()`. Passamos por parâmetro a lista de produtos.

### Usando UUIDs
Outra forma de criar IDs - bastante utilizada no desenvolvimento de software - é usando o UUID (Universally Unique IDentifier)

Essa abordagem utiliza uma biblioteca externa, que precisa ser instalada para poder ser usada no código. Mais explicações e instruções de como fazer a instalação [aqui](https://github.com/rafack/digport-academy/blob/main/logica-programacao/pacotes.md#importando-pacotes-de-projetos-externos)

- Criar uma função `NewUUID` em um arquivo `id.go` (sugiro que fique dentro do pacote model)
- Para gerar um novo uuid, podemos usar [essa](https://pkg.go.dev/github.com/google/uuid#New) função. Observe que o tipo do retorno dessa função é do tipo `UUID`
- [Esse](https://pkg.go.dev/github.com/google/uuid#UUID.String) método retorna uma `string`, que podemos usar como retorno da nossa função
- `ID` do `produtoNovo` `=` `NewUUID()`

[Sobre UUID](https://medium.com/trainingcenter/o-que-%C3%A9-uuid-porque-us%C3%A1-lo-ad7a66644a2b). Diria que o mais interessante é entender o porquê usar ele em vez de IDs sequenciais.
