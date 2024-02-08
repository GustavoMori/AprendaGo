// Utilizando a solução do exercício anterior:
//	1. Em package-level scope, atribua os seguintes valores às variáveis:
// 		1. para "x" atribua 42
// 		2. para "y" atribua "James Bond"
// 		3. para "z" atribua true
//	2. Na funçao main:
//		1. Use fmt.Srpintf para atribuir todos esses valores a uma única variável. Faça essa atribuiçao de tipo string a uma variavel de nome "s" utilizando
//			o operador curto
//		2. Demonstre a variável "s"

package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v %v %v", x, y, z)

	fmt.Println(s)
}
