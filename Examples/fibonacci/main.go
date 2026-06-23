package main

import "fmt"

func fibRec(n int) int {
	if n <= 1 {
		return n
	}

	return fibRec(n-1) + fibRec(n-2)
}

func fibClosure() func() int { // defino que a minha funcao retorna outra funcao
	n1, n2 := 0, 1 // seto os valores inicias que vao ser referenciados em todas as outras execucoes da funcao retornada (ponteiro)
	return func() int {
		fib := n1          // pego o valor a retornar
		n1, n2 = n2, n1+n2 // faco a conta incremental nos valores de referencia guardados em heap, go analisa o lado direito inteiro antes de atribuir
		return fib         // retorno o atual
	}
}

func main() {
	for i := range 10 {
		fmt.Println(fibRec(i))
	} // fibonacci padrao por recursao

	fmt.Println()

	f2 := fibClosure() // faco a chamada a funcao, que seta os valores inicias e retorna a closure a ser executada pela variavel f2
	for range 10 {
		fmt.Println(f2()) // executo a closure incrementalmente
	}
}
