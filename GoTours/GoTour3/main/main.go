package main

import (
	"fmt"
	"math"
)

// --- Tipos ---

type Pessoa struct {
	nome  string
	idade int
}

type Vertex struct {
	Lat, Long float64
}

// --- Variaveis de pacote ---

var m map[string]Vertex // map nil, inicializado depois com make

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

// literal de map: a chave "Google" omite o nome do tipo Vertex
var mapa = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": {
		37.42202, -122.08408,
	},
}

// --- Funcoes auxiliares ---

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// compute recebe uma funcao como argumento e a aplica em (3, 4)
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	// --- Ponteiros ---
	i, j := 20, 11
	var k, n int
	var p *int
	p2 := &i
	p = &j

	fmt.Println(k, float64(n)) // valores zero + conversao de tipo
	fmt.Println(*p)            // dereferencia: valor de j
	fmt.Println(*p2)           // dereferencia: valor de i

	// --- Structs ---
	fmt.Println(Pessoa{"Joaquim", 21})

	var pessoa1 Pessoa // valor zero: {"" 0}
	pessoa2 := Pessoa{"rafael", 22}
	fmt.Println(pessoa1, pessoa2)

	pessoa1.nome = "Joca"
	fmt.Printf("Os nomes sao %s e %s\n", pessoa1.nome, pessoa2.nome)

	p3 := &pessoa2
	p3.idade = 55 // acesso via ponteiro (Go dereferencia sozinho)
	fmt.Println("A idade da pessoa", p3.nome, "virou", p3.idade)

	// --- Arrays ---
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// --- Slice a partir de um array ---
	primos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primos)

	var slicep []int = primos[1:4]
	fmt.Println(slicep)
	fmt.Println()

	// --- Slices compartilham o array subjacente ---
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a1 := names[0:2]
	b := names[1:3]
	fmt.Println(a1, b)

	b[0] = "XXX" // altera tambem 'names' e 'a1' (mesmo array por baixo)
	fmt.Println(a1, b)
	fmt.Println(names)

	// --- Slice literal ---
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	// --- Reslicing (len e cap) ---
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	fmt.Printf("Capacidade eh %d e o tamanho eh %d\n", cap(s), len(s))

	s = s[:4]
	fmt.Printf("Capacidade eh %d e o tamanho eh %d\n", cap(s), len(s))

	// --- Slice nil ---
	var sling []int
	fmt.Println(sling, len(sling), cap(sling))
	fmt.Println("nil!")

	// --- make (slices dinamicos) ---
	makeSlice := make([]int, 5)
	printSlice("makeSlice", makeSlice)

	makeSlice2 := make([]int, 0, 5)
	printSlice("makeSlice2", makeSlice2)

	makeSlice3 := makeSlice2[:2]
	printSlice("makeSlice3", makeSlice3)

	makeSlice4 := makeSlice3[2:5]
	printSlice("makeSlice4", makeSlice4)

	// --- append ---
	makeSlice = append(makeSlice, 1)
	printSlice("makeSlice", makeSlice)

	// --- range ---
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// --- Maps ---
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
	fmt.Println(mapa)

	// --- Funcoes como valores / closures ---
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))

	fmt.Println(compute(hypot))    // passa a closure
	fmt.Println(compute(math.Pow)) // passa uma funcao da stdlib
}
