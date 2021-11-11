package main

import (
	"fmt"
)

func printBoard2(p []string) {
	gridNbr := 1
	pos := 0
	for i := 0; i < 11; i++ {
		for j := 0; j < 11; j++ {
			if (j == 0 || j == 4 || j == 8) && (i == 0 || i == 4 || i == 8) {
				fmt.Printf("%d", gridNbr)
				gridNbr++
			} else if (j == 1 || j == 5 || j == 9) && (i == 1 || i == 5 || i == 9) {
				if p[pos] == "X" {
					fmt.Printf("%s", "X") // These are like this now in order to facilate printing colors
				} else if p[pos] == "O" {
					fmt.Printf("%s", "O")
				} else {
					fmt.Printf("%s", p[pos])
				}
				pos++
			} else if j == 3 || j == 7 {
				fmt.Printf("|")
			} else if i == 3 || i == 7 {
				fmt.Printf("-")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
