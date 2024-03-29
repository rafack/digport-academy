package main

import "welcome"

func main(){
  welcome.sayHello() // Não funciona! pois a função não é exportada / é protegida (inicia com letra minúscula)
  welcome.SayGoodbye()
}
