package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type moves []string

func startGame() {
	selection := true
	fmt.Println("Ready? (y/n)")
	for selection {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		answer := scanner.Text()
		if answer == "y" {
			playGame()
			os.Exit(1)
		} else if answer == "n" {
			fmt.Println("Ok, exiting game...")
			os.Exit(1)
		} else {
			fmt.Println("Expecting either 'y' for yes or 'n' for no. 'n' will exit game")
			fmt.Println("Ready? (y/n)")
		}
	}
}

func playGame() {
	activeGame := true
	moveCount := 0
	player := 1
	tiles := moves{"_", "_", "_", "_", "_", "_", "_", "_", "_"}
	for activeGame {
		if player == 3 {
			player = 1
		}
		printBoard2(tiles)
		if player == 1 {
			tiles[tiles.getPlayerMove(player)] = "X"
			moveCount++
		} else {
			tiles[tiles.getPlayerMove(player)] = "O"
			moveCount++
		}
		if tiles.checkGameStatus(player) {
			printBoard2(tiles)
			fmt.Printf("Player %d wins! Game over\n", player)
			activeGame = false
		} else if moveCount == 9 && activeGame {
			printBoard2(tiles)
			fmt.Printf("Stalemate!\n")
			os.Exit(1)
		}
		player++
	}
}

func (m moves) getPlayerMove(p int) int {
	validInput := false
	for !validInput {
		fmt.Printf("Player %d, please enter a free tile between 1-9: ", p)
		var input string
		i, err := fmt.Scanln(&input)
		if i != 1 || err != nil {
			fmt.Printf("Error, please provide only a single number! ")
		} else {
			nbr, err := strconv.Atoi(input)
			if (nbr < 0 || nbr > 9) || err != nil {
				fmt.Printf("Error, please enter a valid number! ")
			} else {
				nbr-- // moves slice is from 0-8 but players see and enter 1-9
				if !m.validateMove(nbr) {
					fmt.Printf("Error, tile already occupied! ")
				} else {
					return nbr
				}
			}
		}
	}
	return 0 // Never supposed to hit this anyways
}

func (m moves) validateMove(grid int) bool {
	return m[grid] == "_"
}

func (m moves) checkGameStatus(p int) bool {
	var char string
	if p == 1 {
		char = "X"
	} else {
		char = "O"
	}
	return (m.checkHorizontal(char) || m.checkVertical(char) || m.checkDiagonal(char))
}

func (m moves) toString() string {
	return strings.Join([]string(m), "")
}

func (m moves) checkHorizontal(token string) bool {
	tokens := ""
	for i := 0; i < 3; i++ {
		tokens += token
	}
	return (m[:3].toString() == tokens || m[3:6].toString() == tokens || m[6:].toString() == tokens)
}

func (m moves) checkVertical(token string) bool {
	tokens := ""
	for i := 0; i < 3; i++ {
		tokens += token
	}
	for i := 0; i < 3; i++ {
		verticalString := ""
		for j := i; j < 9; j += 3 {
			verticalString += m[j]
		}
		if verticalString == tokens {
			return true
		}
	}
	return false
}

func (m moves) checkDiagonal(token string) bool {
	tokens := ""
	for i := 0; i < 3; i++ {
		tokens += token
	}
	diagonal := ""
	for i := 0; i < 9; i += 4 {
		diagonal += m[i]
	}
	if diagonal == tokens {
		return true
	}
	diagonal = ""
	for i := 2; i < 7; i += 2 {
		diagonal += m[i]
	}
	return diagonal == tokens
}
