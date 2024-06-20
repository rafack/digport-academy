# Lógica de Programação

### 1. Planejamento de Finanças
Crie um programa para ajudar a planejar as finanças.

**Passos:**
1. Crie a sua lista de despesas a serem pagas (ex: `[ mercado, aluguel, eletricidade ]` etc).
2. Use um loop para imprimir todas as despesas.

   A Tour of Go: https://go.dev/tour/moretypes/16 (Range)
   e https://go.dev/tour/moretypes/17 (continuação)
   
   **Recomendação:** olhar o código do Tour of Go, ir digitando e testando, ver a especificação da linguagem https://go.dev/ref/spec e documentação https://go.dev/doc/
    
4. Utilize condicionais para verificar se um item específico (ex: `condomínio`) está na lista e imprima uma mensagem apropriada. Receba esse item pelo teclado.
5. Formate uma string para exibir uma mensagem com o total de itens na lista.

### 2. Controle de Gastos
Desenvolva um programa para calcular o total de gastos mensais.

**Passos:**
1. Crie variáveis para armazenar os gastos fixos (ex: `aluguel`, `internet`, `transporte`).
2. Calcule o total de gastos somando todas as variáveis.
3. Use uma condicional para verificar se o total de gastos excede um determinado orçamento (ex: R$2.000) e imprima uma mensagem apropriada.

## Arrays, Slices, Structs, Maps

### 3. Lista de Contatos 

**Descrição:** Crie um programa para gerenciar uma lista de contatos.

**Passos:**
1. Defina um `struct` para representar um contato, com campos para `nome`, `telefone` e `email`.
2. Crie um slice para armazenar múltiplos contatos.
3. Adicione alguns contatos ao slice.
4. Use um loop para imprimir todos os contatos na lista.

### 4. Map de Contatos 
1. Crie um map com a chave `string` nome e o value com uma `struct` de contatos
2. Informe um contato usando a função scan, e verifique se o contato existe no map criado
3. Imprima os detalhes do contato encontrado, caso não encontrado imprima informando que o contato de "nome" não foi encontrado.

# Git, Go Module, HTTP e API

### 5. Criando uma API Simples 

**Descrição:** Crie uma API em Go para gerenciar uma lista de contatos.
1. Defina um `slice` e adicione uma `struct` de contatos
2. Crie um método http que liste os contatos previamente criados
3. Para a definição de rotas, sugestão usar gorilla mux: https://aprendagolang.com.br/implementando-uma-api-com-gorilla-mux/ . Para baixar a lib use o comando no terminal: go get -u github.com/gorilla/mux .
