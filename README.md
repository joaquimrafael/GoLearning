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

## Métodos (Methods)

Em Go não tenho suporte nativo a classes, mas posso criar **métodos** em tipos. Um
método é só uma função com um argumento especial, o **receiver** (recebedor), que
aparece na sua própria lista de argumentos, entre a palavra `func` e o nome do método.

```go
type Pessoa struct {
    Nome  string
    Idade int
}

// (p Pessoa) é o receiver
func (p Pessoa) Saudacao() string {
    return "Olá, " + p.Nome
}
```

- Chamo o método a partir de uma variável do tipo criado: `pessoa.Saudacao()`.
- Posso declarar métodos em tipos que **não** são struct também.
- Só consigo declarar um método com um receiver cujo tipo está definido **no mesmo
  pacote** que o método. Não dá para declarar método em tipos de outro pacote (inclui os
  tipos embutidos, como `int`).

### Uso de recebedor de ponteiro ou de valor (Pointer vs value receiver)

Posso declarar métodos com **ponteiro recebedor** (`*T`). Métodos com pointer receiver
conseguem **modificar** o valor para o qual o receiver aponta:

```go
func (p *Pessoa) Aniversario() {
    p.Idade++ // altera a struct original
}
```

- Funções que recebem um **ponteiro** como argumento só aceitam um ponteiro. Já os
  métodos com **pointer receiver** aceitam tanto um valor quanto um ponteiro na chamada
  (o Go ajusta automaticamente).
- O contrário também vale: funções que recebem um **valor** só aceitam aquele tipo de
  valor, mas métodos com **value receiver** aceitam valor ou ponteiro na chamada.

Uso o **ponteiro** quando:

- preciso **modificar** o valor para o qual o receiver aponta, ou
- quero **evitar copiar** o valor a cada chamada (mais eficiente para uma struct com
  vários campos).

> Em C#, isso lembra métodos de instância: o value receiver é como um tipo de valor
> copiado, e o pointer receiver é como trabalhar sobre a referência do objeto original.

## Interfaces (Interfaces)

Uma **interface** é um conjunto de assinaturas de método. Um valor de interface pode
armazenar **qualquer tipo** que implemente todos os métodos dessa interface.

```go
type Saudavel interface {
    Saudacao() string
}
```

- Um tipo implementa uma interface só por **implementar os seus métodos** — não há
  declaração explícita de intenção, não existe palavra-chave `implements`.
- Por serem **implícitas**, a definição da interface fica desacoplada da implementação,
  que pode aparecer em qualquer pacote sem combinação prévia.
- Por baixo dos panos, um valor de interface é uma tupla `(value, type)` — guarda um
  valor de um tipo concreto subjacente. Chamar um método na interface executa o método de
  mesmo nome no tipo concreto que ela guarda.
- Para um mesmo tipo implementar uma interface, devo usar value receiver **ou** pointer
  receiver em **todos** os métodos — não posso alternar, senão a interface não aceita o
  tipo.
- Se o valor concreto dentro da interface for `nil`, o método é chamado com um **nil
  receiver** (não dá panic sozinho). Já uma interface `nil` de verdade não guarda nem
  valor nem tipo concreto.
- Interfaces com **0 métodos** são a interface vazia (`interface{}`) e podem guardar
  valores de qualquer tipo.

> Em C# a interface é parecida, mas lá preciso declarar `: IMinhaInterface`. Em Go a
> implementação é implícita. A `interface{}` lembra o `object`/`dynamic`.

### Afirmações de tipo (Type Assertions)

Uma type assertion dá acesso ao valor concreto subjacente de uma interface:

```go
t := i.(T)        // panic se i não for T
t, ok := i.(T)    // ok = false se não for (sem panic)
```

### Switch de tipos (Type switch)

Construção que permite uma sequência de type assertions de uma vez:

```go
switch v := i.(type) {
case int:
    fmt.Println("int:", v)
case string:
    fmt.Println("string:", v)
default:
    fmt.Println("outro tipo")
}
```

### Stringers

Interface ubíqua definida pelo pacote `fmt`. Um **Stringer** é um tipo que sabe se
descrever como string. O `fmt` (e muitos outros pacotes) procura por essa interface na
hora de imprimir valores:

```go
type Stringer interface {
    String() string
}

func (p Pessoa) String() string {
    return fmt.Sprintf("%s (%d anos)", p.Nome, p.Idade)
}
```

> Em C# isso é exatamente o `ToString()` sobrescrito.

## Erros (Errors)

Programas em Go tratam estado de erro com **valores de `error`**. O tipo `error` é uma
interface embutida, parecida com o `fmt.Stringer`:

```go
type error interface {
    Error() string
}
```

- Funções muitas vezes retornam um `error`, e quem chama deve saber tratá-lo (normalmente
  comparando com `nil`):

```go
if err != nil {
    fmt.Println("deu erro:", err)
    return
}
```

**Exemplo de uso:** uma função que retorna um erro. O jeito mais simples de criar um é
com `errors.New` (ou `fmt.Errorf`, que permite formatar):

```go
import (
    "errors"
    "fmt"
)

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("divisão por zero")
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(10, 0)
    if err != nil {
        fmt.Println("erro:", err) // erro: divisão por zero
        return
    }
    fmt.Println(resultado)
}
```

**Erro customizado:** como `error` é só uma interface, qualquer tipo com o método
`Error() string` vira um erro. Assim posso carregar mais dados junto:

```go
type ErroValidacao struct {
    Campo string
    Msg   string
}

func (e *ErroValidacao) Error() string {
    return fmt.Sprintf("%s: %s", e.Campo, e.Msg)
}

func validarIdade(idade int) error {
    if idade < 0 {
        return &ErroValidacao{Campo: "idade", Msg: "não pode ser negativa"}
    }
    return nil
}
```

> Em C# o normal é lançar **exceções**; em Go o erro é um valor retornado que eu checo na
> mão. O `ErroValidacao` lembra criar uma classe que herda de `Exception`, mas aqui basta
> implementar o método `Error()`.

## Leitores (Readers)

O pacote `io` define a interface `io.Reader`, que representa a ponta de leitura de um
fluxo de dados. A biblioteca padrão tem muitas implementações dela: arquivos, conexões de
rede, compressores, cifras e outros.

A `io.Reader` tem o método `Read`:

```go
func (T) Read(b []byte) (n int, err error)
```

- O `Read` preenche o byte slice dado com dados e retorna o número de bytes preenchidos e
  um valor de erro. Retorna o erro `io.EOF` quando o fluxo termina.

**Exemplo de uso:** o `strings.NewReader` me dá um `io.Reader` a partir de uma string.
Leio em pedaços para um buffer (`make([]byte, 8)`) num loop, até chegar no `io.EOF`:

```go
import (
    "fmt"
    "io"
    "strings"
)

func main() {
    r := strings.NewReader("Olá, Go!")
    buf := make([]byte, 8) // leio de 8 em 8 bytes

    for {
        n, err := r.Read(buf)
        fmt.Printf("li %d bytes: %q\n", n, buf[:n])

        if err == io.EOF { // fim do fluxo
            break
        }
    }
}
```

- Reparo que uso `buf[:n]` (e não `buf` inteiro): só os primeiros `n` bytes foram
  preenchidos nesta leitura; o resto do buffer pode ter lixo da iteração anterior.

> Em C# isso lembra a `Stream` (ex.: `StreamReader`) — leio o fluxo aos poucos para um
> buffer até acabar.

## Imagens (Images)

O pacote `image` define a interface `Image`:

```go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

**Exemplo de uso:** crio uma imagem RGBA com `image.NewRGBA` (ela já implementa a
interface `Image`) e uso os métodos `Bounds()` e `At()` para inspecioná-la:

```go
import (
    "fmt"
    "image"
)

func main() {
    m := image.NewRGBA(image.Rect(0, 0, 100, 100))

    bounds := m.Bounds()
    fmt.Println(bounds)           // (0,0)-(100,100)

    cor := m.At(0, 0).RGBA()      // cor do pixel (0,0): r, g, b, a
    fmt.Println(cor)              // 0 0 0 0 (preto transparente, o valor zero)
}
```

- O `Bounds()` me dá o retângulo que delimita a imagem, e o `At(x, y)` devolve a
  `color.Color` daquele pixel (que tem o método `RGBA()` para extrair os componentes).
  
## Funções genéricas (Generic functions)

Posso escrever funções que funcionam com **múltiplos tipos** usando **parâmetros de
tipo**. Eles aparecem entre **colchetes**, antes da lista de argumentos da função:

```go
func Index[T comparable](s []T, x T) int
```

- Isso significa que `s` é um slice de qualquer tipo `T` que satisfaça a restrição
  embutida `comparable`. O `x` também é um valor do mesmo tipo `T`.
- `comparable` é uma restrição (constraint) útil que permite usar os operadores `==` e
  `!=` em valores daquele tipo. No exemplo, uso isso para comparar um valor com cada
  elemento do slice até achar uma correspondência.
- Assim, essa `Index` funciona para **qualquer tipo** que suporte comparação.

> Em C# isso lembra os **generics** (`int Index<T>(T[] s, T x)`), e a restrição
> `comparable` lembra um `where T : ...` limitando o que o tipo pode fazer.

## Tipos genéricos (Generic types)

Um tipo também pode ser **parametrizado** com um parâmetro de tipo. Isso é útil para
implementar **estruturas de dados genéricas** (que servem para vários tipos):

```go
// uma pilha (stack) que guarda qualquer tipo T
type Stack[T any] struct {
    itens []T
}
```

> Em C# isso equivale a uma classe genérica, como `Stack<T>`.

## Goroutines

Uma **goroutine** é uma thread leve gerenciada pelo runtime do Go. É o que me permite
trabalhar com concorrência nos programas.

- Para iniciar uma nova goroutine, uso a palavra-chave `go` antes da chamada:

```go
go f(x, y, z) // roda f concorrentemente
```

- A avaliação de `f`, `x`, `y` e `z` acontece na goroutine **atual**; só a **execução**
  de `f` é que acontece na **nova** goroutine.
- Todas as goroutines rodam no mesmo espaço de endereçamento, então o acesso à memória
  compartilhada precisa ser **sincronizado**. O pacote `sync` traz primitivas úteis para
  isso (ver `sync.Mutex` abaixo), mas em Go a forma mais idiomática costuma ser usar
  **channels**.

> Em C# isso lembra `Task.Run`/`Thread`, mas a goroutine é muito mais barata — posso ter
> milhares delas.

## Canais (Channels)

Um **channel** é um conduíte **tipado** que permite a comunicação entre goroutines:
mando ou recebo valores com o operador de canal `<-`. Os dados fluem na direção da seta.

```go
ch <- v    // envia v para o canal ch
v := <-ch  // recebe do canal ch e atribui o valor a v
```

- Assim como maps e slices, um canal precisa ser criado antes do uso, com `make`:

```go
ch := make(chan int)
```

- Por padrão, envios e recebimentos **bloqueiam** até o outro lado estar pronto. É isso
  que deixa as goroutines se sincronizarem sem locks ou variáveis de condição explícitas.

### Canais com buffer (Buffered channels)

Canais podem ter **buffer**. Passo o tamanho do buffer como segundo argumento do `make`:

```go
ch := make(chan int, 100)
```

- O envio só bloqueia quando o buffer está **cheio**; o recebimento só bloqueia quando o
  buffer está **vazio**.

### Range e Close

Quem envia pode **fechar** um canal para indicar que não vai mandar mais valores. Quem
recebe consegue testar se o canal foi fechado usando o segundo valor do recebimento:

```go
v, ok := <-ch // ok é false se o canal está fechado e não há mais valores
```

- O laço `for i := range c` recebe valores do canal repetidamente até ele ser **fechado**.
- **Só quem envia** deve fechar o canal, nunca quem recebe. Enviar num canal já fechado
  causa **panic**.
- Canais não são como arquivos: normalmente **não** preciso fechá-los. Fechar só é
  necessário quando quem recebe precisa saber que não vêm mais valores — por exemplo,
  para terminar um laço `range`.

### Select

O `select` me permite esperar **várias** operações de comunicação ao mesmo tempo.

- Ele **bloqueia** até que um dos `case` possa rodar; se mais de um estiver pronto,
  escolhe um **aleatoriamente**.
- O `case default` roda quando nenhum outro `case` está pronto — útil para tentar um
  envio/recebimento **sem bloquear**:

```go
select {
case i := <-c:
    // usa i
default:
    // receber de c bloquearia agora
}
```

## Exclusão mútua (Sync.Mutex)

Um **mutex** (mutual exclusion) é uma estrutura que impede que duas goroutines acessem o
mesmo recurso ao mesmo tempo, evitando **race condition** e **deadlock**.

- A biblioteca padrão oferece o `sync.Mutex`, com dois métodos: `Lock` e `Unlock`.
- Posso usar `defer` para garantir que o mutex seja destravado ao fim da função.

```go
import "sync"

type Contador struct {
    mu sync.Mutex
    n  int
}

func (c *Contador) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock() // destrava ao sair, mesmo se der panic
    c.n++
}
```

> Em C# isso lembra o `lock (obj) { ... }` (que usa um `Monitor` por baixo) ou um
> `Mutex`/`SemaphoreSlim` para proteger uma seção crítica.
