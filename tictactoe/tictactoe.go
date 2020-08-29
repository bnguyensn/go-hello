package tictactoe

import (
	"fmt"
	"strings"
)

// PrintBoard logs a tic-tac-toe board
func PrintBoard(board [][]string) {
	// Classic for loop
	// for i := 0; i < len(board); i++ {
	// 	fmt.Printf("%s\n", strings.Join(board[i], " "))
	// }

	// Range for loop
	for _, row := range board {
		fmt.Printf("%s\n", strings.Join(row, " "))
	}
	fmt.Println("")
}

// MakeAMove marks a tic-tac-toe board with a marker.
func MakeAMove(board [][]string, marker string, x int, y int) {
	board[x][y] = marker

	PrintBoard(board)
}

// CreateBoard creates a tic-tac-toe board of size n
func CreateBoard(n int) [][]string {
	board := make([][]string, n)

	for i := 0; i < n; i++ {
		row := make([]string, n)

		for j := range row {
			row[j] = "_"
		}

		board[i] = row
	}

	PrintBoard(board)

	return board
}
