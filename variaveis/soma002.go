package main

import "fmt"

func main() {
	var n1 int
	var n2 int

	fmt.Println("Digite dois numero")
	fmt.Scanln(&n1)
	fmt.Println("Digite um segundo numero")
	fmt.Scanln(&n2)
	fmt.Println("A soma destes dois numeros é igual à:", n1+n2)
	fmt.Println("A subtração destes dois numeros é igual à:", n1-n2)
	fmt.Println("A divisão destes dois numeros é igual à:", n1/n2)
	fmt.Println("A multiplicação destes dois numeros é igual à:", n1*n2)
}
