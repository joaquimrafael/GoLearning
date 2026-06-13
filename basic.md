# Go Guide: From Beginning

## Comandos

### Compile

```sh
go build program.go
```

### Run

```sh
./executable
```

### Compile and Run

```sh
go run program.go
```

### Init mod

```sh
go mod init "module_name"
```

### Format Code

```sh
go fmt program.go
go fmt ./...   # todos os arquivos do diretório
```

### Search non syntax bugs

```sh
go vet program.go
go vet ./...   # todos os arquivos do diretório
```

## Tipos

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

## Funções

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

## Variáveis

`var x tipo` declara o tipo da variável. Posso declarar várias juntas:

```go
var nome, sobrenome, amigo string
```

Posso declarar atribuindo o valor também (dispensa o tipo):

```go
var c, python, java = true, false, "no!"
```

- `:=` cria e atribui a variável com o tipo da atribuição (short declaration) — só dentro da função.
- `=` só atribui.

## Constantes

Declaro o tipo como nas variáveis, mas **não** posso usar `:=`.

```go
const pi = 3.14
```

Constantes numéricas são valores de alta precisão. Constantes não tipadas assumem o tipo necessário pelo contexto.

## Loops

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

## Condicionais

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

## Defers

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

## Pointers

Armazena o endereco de memoria de um valor
*T e o ponteiro para T, valor zerado e nil
& gera um ponteiro
"*"seta o valor interno do ponteiro
Nao tem aritmetica de ponteiro
