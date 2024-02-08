// Utilizando a solucao do exercicio anterior:
// 	1. Em package-level scope, utilizando a palavra-chave var, crie uma variavel com o identificador "y".
//		 O tipo desta variavel deve ser o tipo subjacente do tipo que voce criou no exercicio anterior.
//	2. Na funcao main:
// 		1. Isto já deve estar feito:
//			1. Demonstre o valor da variavel "x"
//			2. Demonstre o tipo da variavel "x"
//			3. Atribua 42 a variavel "x" utilizando o operador "="
//			4. Demonstre o valor da variavel "x"
//		2. Agora faca tambem
//			1. Utilize conversao para transformar o tipo do valor da variavel "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
//			2. Demonstre o valor da variavel "y"
//			3. Demonstre o tipo da variavel "y"

package main

import "fmt"

type xablau int

var x xablau
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
