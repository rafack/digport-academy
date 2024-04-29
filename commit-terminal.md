# Fazendo um commit pelo terminal

Podemos começar com o comando `git status` para verficar as alterações que foram feitas

Usamos `git add` para ir adicionando os arquivos que queremos enviar para nosso GitHub. Podemos fazer isso de um em um arquivo: quero adicionar o arquivo `main.go`, então executo `git add main.go`. 
Depois, adiciono `README.md` com `git add README.md` e assim por diante. Isso também é útil caso eu tenha modificado os dois arquivos (`main.go` e `README.md`) mas só queira adicionar um deles ao meu commit.
Caso eu queira adicionar todos os arquivos que foram modificados ao meu novo commit, uso `git add .`.

Podemos usar `git status` novamente para verificar quais foram os arquivos adicionados

A seguir, utilizamos o comando `git commit -m "<mensagem-do-commit>"` para criar o nosso commit. Substituímos `<mensagem-do-commit>` por uma mensagem descritiva, que diz o que são as alterações que estão sendo adicionadas

Por fim, podemos dar `git status` novamente, só para conferir que está tudo certo. Nesse momento, a resposta do comando status deve ser algo como:
```shell
Your branch is ahead of 'origin/main' by 1 commit.
```

O que significa que o código que temos no PC (costumamos chamar de LOCAL) está na frente do código que está no GitHub (também chamamos de REMOTO ou ORIGIN). 
Nosso código local está à frente do remoto por UM COMMIT. Que é justamente o commit que acabamos de criar no passo anterior.

Podemos confirmar essa afirmação da seguinte forma:
- ir no GITHUB e abrir o repositório que está sendo alterado
- ver o histórico de commits clicando aqui:

 ![image](https://github.com/rafack/digport-academy/assets/70387077/40445e56-bbbd-4028-af74-988e36ea266a)

Observe quais são os commits que estão no REMOTO

Agora, vamos comparar com o que está no LOCAL, usando o comando `git log`. Com esse comando vemos o histórico de commits do nosso local:

O commit que aparece no topo é o mais recente

<img width="654" alt="image" src="https://github.com/rafack/digport-academy/assets/70387077/7e5c4438-777b-4d06-81e2-5816da14a32a">


 ✅ Confirmamos que o nosso LOCAL está à frente do REMOTO por 1 commit! 🙌

Com isso, podemos finalmente executar `git push` 🏁 e ver nosso novo commit no remoto 🥳
