 # JWT - Json Web Tokens 
 - Documentação: https://pkg.go.dev/github.com/golang-jwt/jwt/v5@v5.2.1
 - Debugger: https://jwt.io/#debugger-io 
 - Repo: https://github.com/golang-jwt/jwt 
 - Referências:
   - https://medium.com/@cheickzida/golang-implementing-jwt-token-authentication-bba9bfd84d60
   - https://github.com/pinbar/go-mux-jwt/blob/master/README.md

Para testar, crie um request, método POST, chamando a rota de login com um usuario criado na tabela USUARIO, username será o email: 

![Screenshot 2024-10-17 at 19 12 20](https://github.com/user-attachments/assets/3438f624-b944-4aa3-8fbb-27e49f0ec9e4)


Use o token gerado para a authenticação na API de buscar Produtos por nome, passando na aba Headers a chave "Authorization" e no valor "Bearer tokengerado" :
![Screenshot 2024-10-17 at 19 17 13](https://github.com/user-attachments/assets/ea5bc177-41a2-4091-b6f2-8078c798680d)

Ou passe somente o token na aba Auth:

![Screenshot 2024-10-17 at 19 21 37](https://github.com/user-attachments/assets/436da255-166e-4efa-af19-e689076882ac)

## Código da aula:
https://github.com/emigoulart/digport-academy/tree/main

