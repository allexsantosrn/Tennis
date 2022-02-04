# Tennis

## Autor:
- Alexandre Dantas dos Santos

## Objetivos
- Simulação de uma partida de tenis utilizando um recursos de programação concorrente da linguagem de programação Golang.

## Metodologia
- O projeto foi desenvolvido na Linguagem de Programação Golang, através da IDE Visual Studio Code.
- O projeto foi desenvolvido em dois arquivos, denominados tennis1.go e tennis2.go. No primeiro arquivo temos a versão mais simples (básica) de execução da partida de tênis. 
Já no segundo arquivo temos o projeto executado com sua versão extra/avançada.

## Execução
- Para iniciar a aplicação executar os arquivos presentes na pasta raiz do projeto.
- Para a versão 1 do projeto, informar o argumento de pontuação máxima para o vencedor da partida (go run tennis1.go {pontuação máxima}). Exemplo: **go run tennis1go.1 3**. Neste caso, a pontuação máxima para algum jogador vencer a partida será de 3 pontos. 
- Para a versão 2 do projeto, informar os argumentos: O número de pontos para se vencer um game; O número de games para se vencer um set; E o número de sets máximo existente para a partida (go run tennis2.go {pontuação para vencer o game - pontuação para se vencer um set - número de sets máximo da partida }). Exemplo: **go run tennis1go.2 4 6 5**. Neste caso, temos um exemplo de um cenário real de uma partida de tenis. Vence um game o jogador que adquirir 4 pontos(com dois pontos de vantagem). Vence um set, o jogador que atingir 6 games primeiro(com dois sets de vantagem). E vence a partida o jogador que chegar a 3 sets conquistados, com um limite máximo de 5 sets para a partida.

## Informações complementares
- Na primeira versão, temos a implementação de uma partida de tênis simples, onde vence a partida o jogador que alcançar primeiro a pontuação limite estabelecida. Neste caso, admite-se que a partida tem um único set, com apenas um game, onde é declarado vencedor o jogador que atinge a pontuação limite estabelecida.
- Já na segunda versão, temos a implementação de um jogo mais complexo de tenis. Neste cenário é possível definir o número de pontuações para se vencer um game, o número de games para se vencer um set, e o número de sets máximos permitidos na partida de tenis. Para vencer cada game é necessário ter dois pontos de vantagem para o vencedor. O mesmo pode ser dito para se vencer um set (A partida não contém as regras específicas para tie-break). Neste ponto, é necessária a vantagem de dois games para se vencer um set. É declarado vencedor o jogador que atingir a marca de 3 sets conquistados, contanto que haja vantagem de dois sets para este vencedor. Caso falte apenas um set para se atingir o limite máximo de sets, a regra de vantagem de 2 sets para o vencedor não é necessária.
- Em ambas as versões existe o trecho de código: *time.Sleep(1 * time.Second)* que pode ser descomentado para se enxergar a partida de forma mais lenta. 


