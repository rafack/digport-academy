package welcome

import "fmt"

func PresentDigPortAcademy() {
	sayHello() // essa função pode ser chamada aqui porque é do mesmo pacote
	fmt.Println("This is {Digport} Academy")
}
