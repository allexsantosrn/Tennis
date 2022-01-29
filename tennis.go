package main

import (
	"fmt"
	"math/rand"
	"time"
)

var jogador1 = "Nadal"

var jogador2 = "Federer"

func jogador(nome string, turn chan int) {

	for {

		posse, controle := <-turn

		// Com o canal fechado, exibe o vencedor da partida.
		if controle == false {
			fmt.Printf("Jogador %s venceu a partida.", nome)
			fmt.Print("\n")
			return
		}

		num := gerarNumeroAleatorio()

		if num%7 == 0 {

			fmt.Printf("Jogador %s não acertou a bola (Jogada %d).", nome, posse)
			fmt.Print("\n")

			close(turn)
			return
		}

		fmt.Printf("Jogador %s acertou a bola (Jogada %d).", nome, posse)
		fmt.Print("\n")

		posse++
		turn <- posse
	}
}

func main() {

	turn := make(chan int)

	go jogador(jogador1, turn)
	go jogador(jogador2, turn)

	turn <- 1

	time.Sleep(1e9)
}

// Gera um número aleatório até o valor de 50.
func gerarNumeroAleatorio() int {

	numeroAleatorio := rand.Intn(50)
	return numeroAleatorio
}
