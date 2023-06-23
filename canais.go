package main

import (
	"fmt"
	"time"
)

func ping(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "ping" // usado para enviar e receber mensagem pelo canal
	}
}

func pong(c chan string) {
	for i := 0; i < 10; i++ {
		c <- "pong" // usado para enviar e receber mensagem pelo canal
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string) // inicia a variavel do canal

	go ping(c)
	go pong(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada) // scanln imput dados na tela
}
