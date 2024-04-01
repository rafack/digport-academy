package main

import (
	"fmt"

	cpfValidator "github.com/Nhanderu/brdoc"
)

func main() {
	cpf := "93939393939"
	fmt.Println("CPF é válido?", cpfValidator.IsCPF(cpf))
}
