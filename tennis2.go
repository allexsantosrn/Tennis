package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var numMaxSets = 5

var numSets = 3

var numGames = 6

var pointsPerGames = 4

// Jogador 1
var jogador1 = "Nadal"

// Representa o número de sets ganhos do jogador 1.
var numeroSet1 = 0

// Representa o número de games ganhos do jogador 1.
var numeroGames1 = 0

// Representa o número de pontos no game do jogador 1.
var pontosGame1 = 0

// Jogador 2
var jogador2 = "Federer"

// Representa o número de sets ganhos do jogador 2.
var numeroSet2 = 0

// Representa o número de games ganhos do jogador 1.
var numeroGames2 = 0

// Representa o número de pontos no game do jogador 2.
var pontosGame2 = 0

var wg sync.WaitGroup

func jogador(nome string, turn chan int) {

	defer wg.Done()

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

		if num%17 == 0 {

			// Exibe a mensagem de que o jogador não acertou a bola.
			fmt.Printf("Jogador %s não acertou a bola (Jogada %d).", nome, posse)
			fmt.Print("\n")

			// Incrementa pontuação para o jogador 2, caso o jogador 1 não tenha acertado a bola.
			if nome == jogador1 {

				pontosGame2++
				fmt.Printf("Ponto para %s, totalizando %d ponto(s) no game.", jogador2, pontosGame2)
				fmt.Print("\n \n")
				fmt.Printf("Placar atual no game - %s: %d x %d: %s", jogador1, pontosGame1, pontosGame2, jogador2)
				fmt.Print("\n \n")

				// Incrementa pontuação para o jogador 2, caso o jogador 1 não tenha acertado a bola.
			} else {

				pontosGame1++
				fmt.Printf("Ponto para %s, totalizando %d ponto(s) no game.", jogador1, pontosGame1)
				fmt.Print("\n \n")
				fmt.Printf("Placar atual no game  - %s: %d x %d: %s", jogador1, pontosGame1, pontosGame2, jogador2)
				fmt.Print("\n \n")
			}

			// Encerra o game caso alguns dos jogadores atinga o valor de pontuação limite.
			if pontosGame1 >= pointsPerGames {

				if pontosGame1-pontosGame2 >= 2 {
					numeroGames1++
					fmt.Printf("Jogador %s venceu o game.", jogador1)
					fmt.Print("\n \n")
					fmt.Printf("Placar atual no set - %s: %d x %d: %s", jogador1, numeroGames1, numeroGames2, jogador2)
					fmt.Print("\n \n")
					pontosGame1 = 0
					pontosGame2 = 0
				}
			}

			if pontosGame2 >= pointsPerGames {

				if pontosGame2-pontosGame1 >= 2 {
					numeroGames2++
					fmt.Printf("Jogador %s venceu o game.", jogador2)
					fmt.Print("\n \n")
					fmt.Printf("Placar atual no set - %s: %d x %d: %s", jogador1, numeroGames1, numeroGames2, jogador2)
					fmt.Print("\n \n")
					pontosGame2 = 0
					pontosGame1 = 0
				}
			}

			if numeroGames1 >= numGames {

				if numeroGames1-numeroGames2 >= 2 {
					numeroSet1++
					fmt.Printf("Jogador %s venceu o set.", jogador1)
					fmt.Print("\n \n")
					fmt.Printf("Placar parcial em sets - %s: %d x %d: %s", jogador1, numeroSet1, numeroSet2, jogador2)
					fmt.Print("\n \n")
					numeroGames1 = 0
					numeroGames2 = 0
				}
			}

			if numeroGames2 >= numGames {

				if numeroGames2-numeroGames1 >= 2 {
					numeroSet2++
					fmt.Printf("Jogador %s venceu o set.", jogador2)
					fmt.Print("\n \n")
					fmt.Printf("Placar parcial em sets - %s: %d x %d: %s", jogador1, numeroSet1, numeroSet2, jogador2)
					fmt.Print("\n \n")
					numeroGames2 = 0
					numeroGames1 = 0
				}
			}

			if numeroSet1 >= numSets {

				if numeroSet1-numeroSet2 >= 2 {
					fmt.Printf("Jogador %s perdeu a partida.", jogador2)
					fmt.Print("\n")

					close(turn)
					return
				}

				if numeroSet1 == numMaxSets {
					fmt.Printf("Jogador %s perdeu a partida.", jogador2)
					fmt.Print("\n")

					close(turn)
					return
				}

			}

			if numeroSet2 >= numSets {
				if numeroSet2-numeroSet1 >= 2 {
					fmt.Printf("Jogador %s perdeu a partida.", jogador1)
					fmt.Print("\n")

					close(turn)
					return
				}

				if numeroSet2 == numMaxSets {
					fmt.Printf("Jogador %s perdeu a partida.", jogador2)
					fmt.Print("\n")

					close(turn)
					return
				}
			}

			posse++
			turn <- posse

		} else {
			// Exibe a mensagem de acerto na bola.
			fmt.Printf("Jogador %s acertou a bola (Jogada %d).", nome, posse)
			fmt.Print("\n")

			posse++
			turn <- posse
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	//pointsMax, err1 := strconv.Atoi(os.Args[1])

	/*if err1 != nil {
		fmt.Println("Erro!!")
		return
	}

	// Atribui o valor passado via argumento como pontuação máxima para vitória.
	pontosFixos = pointsMax */

	// Com o canal fechado, exibe o vencedor da partida.
	turn := make(chan int)

	wg.Add(2)

	// Goroutines para os dois jogadores
	go jogador(jogador1, turn)
	go jogador(jogador2, turn)

	// Iniciando na jogada um.
	turn <- 1

	wg.Wait()
}

// Gera um número inteiro aleatório até o limite de 50.
func gerarNumeroAleatorio() int {
	numeroAleatorio := rand.Intn(50)
	return numeroAleatorio
}
