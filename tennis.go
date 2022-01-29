package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Representa o número de pontos máximos que identificará o vencedor da partida.
var pontosFixos = 5

// Jogador 1
var jogador1 = "Nadal"

// Representa o número de pontos do jogador 1.
var pontosj1 = 0

// Jogador 2
var jogador2 = "Federer"

// Representa o número de pontos do jogador 2.
var pontosj2 = 0

func jogador(nome string, turn chan int) {

	for {

		posse, controle := <-turn

		// Com o canal fechado, exibe o vencedor da partida.
		if controle == false {
			fmt.Printf("Jogador %s venceu a partida.", nome)
			fmt.Print("\n")
			return
		}

		// Gera um número aleatório que auxilia no processo de não acerto da bola.
		num := gerarNumeroAleatorio()

		if num%7 == 0 {

			fmt.Printf("Jogador %s não acertou a bola (Jogada %d).", nome, posse)
			fmt.Print("\n")

			// Incrementa pontuação para o jogador 2, caso o jogador 1 não tenha acertado a bola.
			if nome == jogador1 {

				pontosj2++
				fmt.Printf("Ponto para %s, totalizando %d ponto(s).", jogador2, pontosj2)
				fmt.Print("\n \n")
				fmt.Printf("Placar atual - %s: %d x %d: %s", jogador1, pontosj1, pontosj2, jogador2)
				fmt.Print("\n \n")

				// Incrementa pontuação para o jogador 2, caso o jogador 1 não tenha acertado a bola.
			} else {

				pontosj1++
				fmt.Printf("Ponto para %s, totalizando %d ponto(s).", jogador1, pontosj1)
				fmt.Print("\n \n")
				fmt.Printf("Placar atual - %s: %d x %d: %s", jogador1, pontosj1, pontosj2, jogador2)
				fmt.Print("\n \n")
			}

			// Encerra a partida caso alguns dos jogadores atinga o valor limite.
			if pontosj1 == pontosFixos || pontosj2 == pontosFixos {
				fmt.Printf("Jogador %s perdeu a partida.", nome)
				fmt.Print("\n")
				close(turn)
				return
			}

		}

		fmt.Printf("Jogador %s acertou a bola (Jogada %d).", nome, posse)
		fmt.Print("\n")

		posse++
		turn <- posse
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	turn := make(chan int)

	go jogador(jogador1, turn)
	go jogador(jogador2, turn)

	// Iniciando na jogada um.
	turn <- 1

	time.Sleep(1e9)
}

// Gera um número inteiro aleatório até o limite de 50.
func gerarNumeroAleatorio() int {

	numeroAleatorio := rand.Intn(50)
	return numeroAleatorio
}
