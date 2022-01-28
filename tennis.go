package main

import (
	"fmt"
	"math/rand"
	"time"
)

var jogador1 = "Nadal"

var jogador2 = "Federer"

func player(nome string, turn chan int) {

	teste := <-turn

	fmt.Println("Teste:", teste)

	turn <- teste
}

func main() {

	turn := make(chan int)

	go player(jogador1, turn)
	go player(jogador2, turn)

	turn <- 1

	time.Sleep(1e9)
}

func gerarNumeroAleatorio() int {

	numeroAleatorio := rand.Intn(50)
	return numeroAleatorio
}
