package main

import "fmt"

func main() {

	resultado := multiplicador(2, 56)
	fmt.Printf("O resultado da multiplicação foi: %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("O resultado da soma foi: %d\r\n", resultado)
}

func multiplicador(x int, y int) int {
	return x * y
}
