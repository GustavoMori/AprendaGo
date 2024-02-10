// Crie um programa que:
// Atribua um valor int a uma variavel
// Demonstre este valor em decimal, binario e hexadecimal
// Desloque os bits dessa variavel 1 pra esquerda e atribua o resultado a outra variavel
// Demonstre esta outra variavel em decimal binario hexadecimal

package main

import "fmt"

func main() {
	x := 23

	fmt.Printf("%d, %#x, %b \n", x, x, x)

	y := x << 1

	fmt.Printf("%d, %#x, %b \n", y, y, y)
}
