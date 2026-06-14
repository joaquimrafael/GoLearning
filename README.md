# Go Guide: From Beginning

## Sobre a linguagem (About the language)

Go (ou Golang) é uma linguagem criada no Google (2009), pensada para ser simples,
rápida de compilar e boa para sistemas concorrentes/servidores.

- **Compilada** para binário nativo (não roda em VM, diferente de C#/Java). Gera um
  executável único, sem runtime externo.
- **Tipagem estática e forte**: o tipo é checado em tempo de compilação e não há
  conversões implícitas (preciso converter na mão, ex.: `float64(x)`).
- **Inferência de tipo**: mesmo sendo estática, o `:=` deduz o tipo pelo valor.
- **Garbage collector**: gerencia a memória pra mim (não dou `free` como em C).
- **Não é orientada a objetos no sentido clássico**: não tem classes nem herança. Tem
  `structs` + `methods` + `interfaces`, e o reuso vem de **composição**, não de herança.
- **Concorrência nativa**: `goroutines` e `channels` fazem parte da linguagem.
- **Opinativa por padrão**: vem com `gofmt` (formatação única), e código com imports ou
  variáveis não usados **não compila**.

> Vindo de **C#**: a sintaxe lembra mais C, mas a experiência (tipagem forte, GC,
> ferramentas integradas) é familiar. As maiores diferenças são a ausência de classes/
> herança e o fato de o binário ser nativo e autossuficiente.

## Comandos (Commands)

### Compilar (Compile)

```sh
go build program.go
```

### Executar (Run)

```sh
./executable
```

### Compilar e executar (Compile and Run)

```sh
go run program.go
```

### Iniciar módulo (Init mod)

```sh
go mod init "module_name"
```

### Formatar código (Format Code)

```sh
go fmt program.go
go fmt ./...   # todos os arquivos do diretório
```

### Procurar bugs não sintáticos (Search non syntax bugs)

```sh
go vet program.go
go vet ./...   # todos os arquivos do diretório
```

## Tipos (Types)

Declaramos os tipos **depois** das variáveis, funções e parâmetros.

```go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias para uint8

rune // alias para int32
     // representa um Unicode code point

float32 float64

complex64 complex128
```

- Se não declarar valor: numéricos são `0`, bool é `false` e string é `""`.
- Consigo transformar o tipo da variável usando `tipo(var)`.

## Funções (Functions)

Parecidas com C: `func nome(params tipo) retorno { codigo }`

- Podem retornar qualquer número de resultados: `(int, string)` / `(string, string)`.
- Posso nomear o retorno também: `split(sum int) (x, y int)` — daí retorno com `return` vazio (naked).

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

## Variáveis (Variables)

`var x tipo` declara o tipo da variável (e ela já nasce com o valor zero do tipo).

Todas as formas possíveis de declarar:

```go
// 1. só o tipo (recebe o valor zero)
var x int

// 2. tipo + valor
var x int = 10

// 3. valor sem tipo (o tipo é inferido pela atribuição)
var x = 10

// 4. várias do mesmo tipo de uma vez
var nome, sobrenome, amigo string

// 5. várias com valores (cada uma infere seu tipo)
var c, python, java = true, false, "no!"

// 6. bloco var (agrupa declarações)
var (
    legal  bool   = true
    contador int  = 0
    msg    string = "oi"
)

// 7. short declaration (:=) — infere o tipo, só DENTRO de função
x := 10
a, b := 1, "dois"
```

- `:=` cria **e** atribui, inferindo o tipo da atribuição (short declaration). Só funciona
  dentro de função e exige pelo menos uma variável nova do lado esquerdo.
- `=` só atribui (a variável já precisa existir).
- O bloco `var ( ... )` é o mesmo que escrever vários `var` seguidos — só organiza melhor.

> Em C#, `var x = 10;` é parecido com o `:=`/`var x = ...` do Go (inferência), e
> declarar o tipo explícito (`int x;`) lembra o `var x int`.

## Constantes (Constants)

Declaro o tipo como nas variáveis, mas **não** posso usar `:=`.

```go
const pi = 3.14
```

Constantes numéricas são valores de alta precisão. Constantes não tipadas assumem o tipo necessário pelo contexto.

## Laços (Loops)

Só existe o `for`, parecido com o do C. Não tenho parênteses entre os componentes:

```go
for i := 0; i < 10; i++ {
    // ...
}
```

- Os três componentes são: `init; condition; post`.
- Variáveis instanciadas no `for` (como o `i`) só são válidas no escopo do loop.
- Os parâmetros de `init` e `post` podem ser implícitos (omitidos).

Também posso dropar os `;` — daí temos o "while" do Go (só o `for` com a condição):

```go
for x < 189 {
    // ...
}
```

Ou seja, quando quero um while, uso o `for` sem o `init` e o `post`. Mas devo tomar
cuidado para evitar loops infinitos (sem condição de parada e/ou atualização).

## Condicionais (Conditionals)

Ifs também têm estrutura parecida — sem parênteses na expressão:

```go
if x > y {
    // ...
}
```

Como no `for`, também posso instanciar uma variável no começo do `if` e usá-la no seu
escopo (só vai ser válida ali). Se tiver um `else`, também posso usá-la lá:

```go
if v := num % 2; v == 0 {
    return true
}
```

O `switch` funciona como nas outras linguagens (if/elses encadeados):

```go
switch dia {
case "sabado":
    fmt.Println("fim de semana")
case "domingo":
    fmt.Println("fim de semana")
default:
    fmt.Println("dia util")
}
```

- Switch sem condicao, if else elegante

```go
switch {
case x < 0:
    fmt.Println("negativo")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positivo")
}
```

- Switch com atribuicao interna

```go
switch nota := calcularNota(); {
case nota >= 7:
    fmt.Println("aprovado")
case nota >= 5:
    fmt.Println("recuperacao")
default:
    fmt.Println("reprovado")
}
```

- O `break` funciona automaticamente (não preciso escrever no fim de cada `case`).
- Os cases não precisam ser constantes, e os valores não precisam ser inteiros.
- `switch` sem condição equivale a `switch true` — uma forma limpa de if-then-else.

## Adiamentos (Defers)

Adia a execução de uma função até que as funções em volta terminem (retornem).

- Os argumentos da chamada são avaliados **imediatamente**, mas a chamada em si só é
  executada quando as funções em volta retornam.
- Essas chamadas adiadas são empilhadas — ou seja, executadas em **LIFO** (a última
  adiada é a primeira a rodar).

```go
defer fmt.Println("mundo")
fmt.Println("olá")
// imprime: olá -> mundo
```

## Ponteiros (Pointers)

Um ponteiro armazena o **endereço de memória** de um valor.

- `*T` é o tipo "ponteiro para `T`". Seu valor zero é `nil`.
- `&` gera um ponteiro (pega o endereço de uma variável).
- `*` na frente do ponteiro acessa/seta o valor interno (dereferência).
- Go **não** tem aritmética de ponteiro (diferente de C).

```go
i := 42
p := &i        // p é *int, aponta para i
fmt.Println(*p) // 42  (lê o valor apontado)
*p = 21         // altera i através do ponteiro
fmt.Println(i)  // 21
```

> Em C# isso lembra a diferença entre tipos por referência e o uso de `ref`/`out`:
> passar um `*T` permite alterar o valor original de quem chamou.

## Estruturas (Structs)

Uma coleção de atributos/campos.

```go
type Pessoa struct {
    Nome  string
    Idade int
}
```

- Instancio passando os parâmetros (na ordem dos campos):

```go
pessoa2 := Pessoa{"rafael", 22}
```

- Acesso os campos com `struct.atributo` (ex.: `pessoa2.Nome`).
- Posso acessar os campos da struct através de **ponteiros** (Go faz a dereferência
  automática: `p.Nome` em vez de `(*p).Nome`).
- Literais de struct representam um valor recém-alocado, listando os valores dos campos.
  Posso nomear só alguns campos (os omitidos ficam com o valor zero):

```go
p2 := Pessoa{Nome: "Joao"} // Idade fica 0
```

- `&` retorna o ponteiro de uma struct: `&Pessoa{...}`.

> Em C# uma `struct` do Go é mais parecida com uma `class`/`record` simples de dados —
> mas sem herança; o reuso em Go vem de composição.

## Vetores (Arrays)

`[n]T` é um vetor de `n` valores do tipo `T`:

```go
var a [10]int
```

- **Não** podem ter o tamanho modificado (o tamanho faz parte do tipo).
- Acesso os elementos pelo índice: `a[1]`.

## Fatias (Slices)

Tamanho dinâmico (os mais comuns). `[]T` é um slice de elementos do tipo `T`.

- Especifico 2 índices, inicial e final: `a[low : high]`.
- Gera um range que **inclui** o primeiro índice mas **exclui** o último — `[ini, end)`:

```go
a[1:4] // -> {a[1], a[2], a[3]}
```

- Um slice **não** guarda dados próprios: ele aponta para parte de um array subjacente.
  Mexer nos elementos do slice mexe no array.
- Slice literal é como um array literal, mas **sem tamanho** (cria o array e referencia nele):

```go
[]bool{true, true, false}
```

- Posso usar defaults nos índices (todos equivalentes quando o slice tem tamanho 10):

```go
a[0:10]  a[:10]  a[0:]  a[:]
```

- Slices têm `len` (nº de elementos) e `cap` (nº de elementos do array subjacente a partir
  do início do slice): `len(s)`, `cap(s)`.
- Valor zero é `nil`.
- Posso criar slices com `make` — é assim que crio vetores dinâmicos em Go:

```go
a := make([]int, 5)    // len(a)=5
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

- Posso criar slices de slices (matrizes).
- Para adicionar valores ao final, uso `append`:

```go
func append(s []T, vs ...T) []T
```

  Se o array subjacente for menor do que o necessário, o Go realoca em outro maior
  automaticamente e o slice passa a apontar para ele.

## Intervalo (Range)

Forma do `for` que itera sobre vetores e slices.

```go
for i, v := range slice {
    fmt.Println(i, v)
}
```

- O range de um slice retorna 2 valores: o **índice** atual e o **elemento** daquele
  índice (uma cópia).
- Posso omitir um dos dois valores, ou pegar só o índice:

```go
for i, _ := range pow      // só o índice (com _)
for _, value := range pow  // só o valor
for i := range pow         // só o índice (sem o segundo)
```

## Mapas (Maps)

Mapeia chaves para valores (hash).

- Valor zero é `nil`.
- `make` retorna um map já inicializado e pronto para uso:

```go
m = make(map[TipoChave]TipoValor)
```

- Literais são como os literais de struct, mas as chaves são obrigatórias. Se o tipo de
  alto nível for apenas um nome de tipo, posso omiti-lo do elemento do literal.

```go
// inserir ou atualizar
m[key] = elem

// ler
elem = m[key]

// deletar
delete(m, key)

// testar se a chave existe (two-value assignment)
elem, ok = m[key]
```

- Em `elem, ok = m[key]`: se a chave está em `m`, `ok` é `true`; senão, `false` (e `elem`
  recebe o valor zero do tipo do valor).
- Se `elem` ou `ok` ainda não foram declarados, posso usar a short declaration:

```go
elem, ok := m[key]
```

## Funções como valores (Functions)

Funções são **valores** também: podem ser passadas como outros valores, usadas como
argumentos de funções e como retornos.

```go
func aplicar(fn func(int) int, x int) int {
    return fn(x)
}
```

- Podem ser **closures**: uma função que referencia variáveis de fora do seu corpo. A
  função consegue ler e atribuir a essas variáveis — nesse sentido ela fica "presa"
  (bound) a elas.

```go
func contador() func() int {
    n := 0
    return func() int {
        n++ // captura e modifica o n de fora
        return n
    }
}
```

> Em C# isso equivale a delegates / `Func<>` e às lambdas que capturam variáveis do escopo.

## Cadeias de caracteres (Strings)

Posso tratar strings de forma parecida com slices e arrays: acesso por índice `s[x]`,
pego o tamanho com `len(s)` e faço slice com `s[ini:fim]`.

```go
s := "golang"
fmt.Println(len(s))   // 6
fmt.Println(s[1:4])   // "ola"
```

- **Atenção:** o índice `s[x]` retorna um `byte` (um `uint8`), **não** um caractere. Por
  isso `len(s)` conta **bytes**, não letras — e em caracteres fora do ASCII (acentos,
  emojis) um único caractere pode ocupar mais de um byte.
- Strings em Go são **imutáveis**: não consigo fazer `s[0] = 'x'`. Para "alterar", crio
  uma nova string.
- É recomendado **evitar fazer slice direto** e preferir as funções do pacote `strings`
  (`strings.Contains`, `strings.Split`, `strings.ToUpper`, `strings.Replace`, etc.), que
  são mais legíveis e seguras:

```go
import "strings"

strings.ToUpper("golang")          // "GOLANG"
strings.Contains("golang", "lan")  // true
strings.Split("a,b,c", ",")        // []string{"a", "b", "c"}
```

> Em C# a `string` também é imutável, e o pacote `strings` faz o papel dos métodos de
> `string`/`System.String` (`ToUpper`, `Contains`, `Split`...). A diferença é que indexar
> uma string em Go te dá um `byte`, não um `char`.

## Conjuntos (Sets)

Não tenho sets nativos em Go. Por isso simulo o funcionamento com um `map`, usando a
**chave** como o tipo dos elementos e o **valor** como `bool`. Assim a chave guarda o
elemento (que é único, pois chaves de map não se repetem) e o `bool` indica que ele está
presente:

```go
set := make(map[string]bool)

// adicionar
set["go"] = true
set["c#"] = true

// remover
delete(set, "c#")

// testar se existe (two-value assignment, como em maps)
if set["go"] {
    fmt.Println("go está no set")
}

// tamanho do set
fmt.Println(len(set)) // 1
```

- Como em qualquer map, ler uma chave que não existe devolve o valor zero do `bool`
  (`false`) — então `set["x"]` já responde se `x` está ou não no set, sem dar erro.
- Um idiomatismo comum é usar `map[T]struct{}` em vez de `bool` quando só me importa a
  presença (não o valor), porque `struct{}` não ocupa memória. Aí testo com o
  `_, ok := set[x]`. Mas começar com `bool` é mais simples de ler.

> Em C# isso equivale ao `HashSet<T>`. Em Go não há um tipo pronto, então o `map` faz
> esse papel.

## Leitura de valores do console (Console input)

Posso ler valores do console de três formas, cada uma com um propósito:

**1. `fmt.Scan` — leio valores separados por espaço**

Passo o **endereço** das variáveis (`&`) e o `Scan` preenche cada uma, separando a
entrada por whitespace (espaço ou quebra de linha):

```go
var nome string
var idade int

fmt.Print("Nome e idade: ")
fmt.Scan(&nome, &idade) // entrada: "joaquim 30"

fmt.Println(nome, idade) // joaquim 30
```

**2. `bufio.Scanner` — leio a linha inteira**

Útil quando o texto tem espaços e eu quero a linha toda (incluindo o que o `Scan`
cortaria no espaço):

```go
scanner := bufio.NewScanner(os.Stdin)
scanner.Scan()           // lê uma linha
linha := scanner.Text()  // a linha já vem sem o '\n'

fmt.Println(linha)
```

**3. `bufio.Reader` — leio até um caractere definido**

Aqui eu escolho o delimitador (ex.: `'\n'`). O `ReadString` devolve tudo até ele,
**incluindo** o próprio delimitador:

```go
reader := bufio.NewReader(os.Stdin)
linha, _ := reader.ReadString('\n') // lê até a quebra de linha

fmt.Print(linha)
```

- O `&` no `fmt.Scan` é obrigatório porque o Go passa argumentos **por valor**; sem o
  endereço, a função não conseguiria alterar a minha variável original.
- O `bufio.Scanner` já tira o `\n`; o `bufio.Reader` (`ReadString`) **mantém** o
  delimitador — se eu não quiser, uso `strings.TrimSpace`.

> Em C# isso lembra o `Console.ReadLine()` (linha inteira, como o `Scanner`) e o
> `Console.Read()` (caractere a caractere). O `fmt.Scan`, lendo por espaços direto nas
> variáveis, parece com um `Console.ReadLine().Split(' ')` já convertido para os tipos.
