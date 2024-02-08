// Crie um tipo. O tipo Subjacente deve ser int
// Crie uma variavel para esse tipo, com o identificador "x", utilizando a palavra chave var
// Na funcao main:
//  1. Demonstre o valor da variavel "x"
//  2. Demonstre o tipo da variavel "x"
//  3. Atribua 42 a variavel "x" utilizando o operador "="
//  4. Demonstre o valor da variavel "x"
package main

import "fmt"

type xablau int

var x xablau

func main() {
	fmt.Println(x)
	fmt.Printf("%T", x)

	x = 42

	fmt.Println(x)
}
