package main

import (
	"fmt"
	"math/rand"
	"time"
)

var jogador1 = "Nadal"

var jogador2 = "Federer"

func jogador(nome string, turn chan int) {

	teste := <-turn

	num := gerarNumeroAleatorio()

	fmt.Println("Teste:", num)

	turn <- teste
}

func main() {

	turn := make(chan int)

	go jogador(jogador1, turn)
	go jogador(jogador2, turn)

	turn <- 1

	time.Sleep(1e9)
}

func gerarNumeroAleatorio() int {

	numeroAleatorio := rand.Intn(50)
	return numeroAleatorio
}
