package main

import "github.com/rafack/digport-academy/welcome"

func main() {
	//welcome.sayHello() Não funciona! pois a função não é exportada / é protegida (inicia com letra minúscula)
	welcome.SayGoodbye()
}
