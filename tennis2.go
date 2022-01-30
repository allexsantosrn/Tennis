package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// Define o limite máximo de sets em uma partida.
var numMaxSets = 5

// Define o número de sets mínimo para garantir uma vitória.
var numSets = 3

// Define o limite de games para vencer um set.
var numGames = 6

// Define o número de pontos para vencer um game.
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

		//time.Sleep(1 * time.Second)

		posse, controle := <-turn //

		// Com o canal fechado, exibe o vencedor da partida.
		if controle == false {
			fmt.Printf("Jogador %s venceu a partida.", nome)
			fmt.Print("\n")
			return
		}

		// Gera um número aleatório que auxilia no processo de não acerto da bola.
		num := gerarNumeroAleatorio()

		// Da faixa de valor aleatório gerado (até 50), verifica se o número gerado é divisível por 7.
		// Caso seja divisível, significa que o jogador não acertou a bola.
		if num%7 == 0 {
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

			// Encerra o game caso o jogador 1 atinja o valor de pontuação limite.
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

			// Encerra o game caso o jogador 2 atinja o valor de pontuação limite.
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

			// Encerra o set caso o jogador 1 atinga o valor de games necessários.
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

			// Encerra o set caso o jogador 1 atinga o valor de games necessários.
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

			//Encerra a partida caso o jogador 1 atinga o valor de sets necessários.
			if numeroSet1 >= numSets {
				if numeroSet1-numeroSet2 >= 2 {
					fmt.Printf("Match point!!!!!!!!!!!!!!!!!")
					fmt.Print("\n")
					fmt.Printf("Jogador %s perdeu a partida.", jogador2)
					fmt.Print("\n")

					close(turn)
					return
				}

				if numeroSet1 == numMaxSets {
					fmt.Printf("Match point!!!!!!!!!!!!!!!!!")
					fmt.Printf("Jogador %s perdeu a partida.", jogador2)
					fmt.Print("\n")

					close(turn)
					return
				}

			}

			//Encerra a partida caso o jogador 2 atinga o valor de sets necessários.
			if numeroSet2 >= numSets {
				if numeroSet2-numeroSet1 >= 2 {
					fmt.Printf("Match point!!!!!!!!!!!!!!!!!")
					fmt.Print("\n")
					fmt.Printf("Jogador %s perdeu a partida.", jogador1)
					fmt.Print("\n")

					close(turn)
					return
				}

				if numeroSet2 == numMaxSets {
					fmt.Printf("Match point!!!!!!!!!!!!!!!!!")
					fmt.Print("\n")
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

			// Incrementa o número de posse de bolas.
			// Repassa a posse para o outro jogador.
			posse++
			turn <- posse
		}
	}
}

func main() {

	rand.Seed(time.Now().UnixNano())

	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("[ERRO]!!! Por favor, informe corretamente os três parâmetros necessários.")
		return
	}

	/* Recebe os valores repassados via argumento e seta o número de pontos por game,
	número de sets por set e o número máximo de sets por partida.  */
	pointsMaxPerGame, err1 := strconv.Atoi(os.Args[1])
	gamesMaxPerSet, err2 := strconv.Atoi(os.Args[2])
	setsMasxperMatch, err3 := strconv.Atoi(os.Args[3])

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("[ERRO]!!")
		return
	}

	// Atribui o valor passado via argumento como pontuação máxima por game.
	pointsPerGames = pointsMaxPerGame

	// Atribui o valor passado via argumento como valor máximo de sets.
	numMaxSets = setsMasxperMatch

	// Atribui o valor passado via argumento como valor máximo de games para vencer um set.
	numGames = gamesMaxPerSet

	// Criando a channel
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
