package main

import "fmt"

var (
	// Endereco é um calor muito importante e tem que ser publico
	Endereco, telefone string // A variável Endereco é publica, já a variavel telefone é privada
	quantidade         int // quantidade = 0
	comprou            bool // comprou = false
	valor              float64 // valor = 0.00
	palavras           rune 
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereco: %s\r\n", Endereco)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("palavras: %v\r\n", palavras)
	fmt.Printf("teste: %v\r\n", teste)
}
