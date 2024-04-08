package main

import (
	"fmt"
	"strconv"
)

func create_empty_board() [3][3]string {
	return [3][3]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}

func main() {
	var board = create_empty_board()
	var turn int = 0
	print_board(board)

	fmt.Printf("It's %d's turn \n", turn)
	for {
		print_board(board)
		if player_turn(turn, board) {
			break
		}
	}

}

func player_turn(turn int, board [3][3]string) bool {
	var col, row int
	fmt.Print("type col then row")
	fmt.Scanf("%v %v", &col, &row)

	if board[row][col] == "_" {
		board[row][col] = strconv.Itoa(turn)
		return true
	} else {
		fmt.Println("Already occupied try again!")
		return false
	}
}
func print_board(board [3][3]string) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			fmt.Printf(board[row][col])
		}
		fmt.Println()
	}
}
