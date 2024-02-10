// Utilizando iota, crie 4 consts cujos valores sejam os proximos 4 anos
// Demonstre esses valores

package main

import "fmt"

const (
	_ = 2024 + iota
	x
	y
	z
	w
)

func main() {
	fmt.Print(x, y, z, w)
}
