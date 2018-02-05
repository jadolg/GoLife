package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
)

const BOARDSIZE = 45
const ITERATIONS  = 20000
const TIME_BETWEEN_ITERATIONS  = time.Second / 10

func printArray(array [BOARDSIZE][BOARDSIZE]int) {
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			if array[i][j] == 0 {
				fmt.Print(" |")
			} else {
				fmt.Print("0|")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func randomlyPopulate(board *[BOARDSIZE][BOARDSIZE]int) {
	fmt.Println("Creating initial population")
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			board[i][j] = random.Intn(2)
		}
	}
}

func liveOrDie(i int, j int, board [BOARDSIZE][BOARDSIZE]int) int {
	c := 0
	for k := i - 1; k < i+2; k++ {
		for l := j - 1; l < j+2; l++ {
			if (k == i && l == j) || k < 0 || l < 0 || k > BOARDSIZE-1 || l > BOARDSIZE-1 {
				continue
			}
			if board[k][l] == 1 {
				c ++
			}
		}
	}
	if c == 3 {
		return 1
	}
	if board[i][j] == 1 && c == 2 {
		return 1
	}
	return 0
}

func step(board [BOARDSIZE][BOARDSIZE]int) [BOARDSIZE][BOARDSIZE]int {
	var aux_board [BOARDSIZE][BOARDSIZE]int
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			aux_board[i][j] = liveOrDie(i, j, board)
		}
	}
	return aux_board
}

func isGameOver(board [BOARDSIZE][BOARDSIZE]int) bool {
	c := 0
	for i := 0; i < BOARDSIZE; i++ {
		for j := 0; j < BOARDSIZE; j++ {
			c += board[i][j]
		}
	}
	return c == 0
}

func main() {
	var board [BOARDSIZE][BOARDSIZE]int
	randomlyPopulate(&board)
	printArray(board)
	for n := 0; n < ITERATIONS; n++ {
		board = step(board)
		printArray(board)
		if isGameOver(board){
			fmt.Printf("Life has ended after %v iterations\n", n)
			os.Exit(1)
		}
		time.Sleep(TIME_BETWEEN_ITERATIONS)
	}
	fmt.Printf("Board survived %v iterations\n", ITERATIONS)
}
