package main

import (
	"fmt"
)

const K float64 = 373.2

func main() {

	T := (K - 273)
	fmt.Printf("O valor convertido é: %G", T)
}
