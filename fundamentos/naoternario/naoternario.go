package main

import "fmt"

// Não tem operador ternário
func obterResultado(nota float64) string {
	// return nota >= 6 ? "Aprovado" : "Reprovado"
	if nota >= 6 {
		return "Aprovado"
	}

	return "Reprovado"
}

func main() {
	fmt.Println("Resultado 6.2 =>", obterResultado(6.2))
	fmt.Println("Resultado 5.999 =>", obterResultado(5.999))
}
