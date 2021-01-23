package main

import "fmt"

func main() {
	var nome, cargo, nomechefe string
	var idade int
	var salario float64
	fmt.Println("Informe seu nome e sua idade:")
	fmt.Scanln(&nome, &idade)
	fmt.Println("Certo, agora informe seu cargo e o nome do seu chefe:")
	fmt.Scanln(&cargo, &nomechefe)
	fmt.Println("E por último, digite seu ultimo salário:")
	fmt.Scanln(&salario)
	fmt.Printf("Os dado informados foram:\nNome: %s\nIdade: %d\nCargo: %s\nNome do chefe: %s\nSalário: %f", nome, idade, cargo, nomechefe, salario)
}
