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
	var turn int = 1

	for {
		fmt.Printf("It's %d's turn \n", turn)
		print_board(board)
		for {
			if player_turn(turn, &board, "sen") {
				break
			}
		}
		if turn == 1 {
			turn = 0
		} else if turn == 0 {
			turn = 1
		}
	}

}

func player_turn(turn int, board *[3][3]string, test string) bool {
	var col, row int
	fmt.Print("type col then row: ")
	fmt.Scan(&col, &row)

	fmt.Printf("%d,%d,%d, %s", col, row, turn, test)
	if board[row][col] == "_" {
		board[row][col] = strconv.Itoa(turn)
		return true
	}
	fmt.Println("Already occupied try again!")
	return false
}
func print_board(board [3][3]string) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			fmt.Printf(board[row][col])
		}
		fmt.Println()
	}
}
