package main

import (
	"fmt"
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
	var turn string = "X"

	for {
		fmt.Printf("It's %s's turn \n", turn)
		print_board(board)
		for {
			if player_turn(turn, &board) {
				break
			}
		}
		if check_win(board) {
			print_board(board)
			fmt.Printf("%s WINS!!!!!\n", turn)
			break
		} else if check_draw(board) {
			print_board(board)
			fmt.Println("You guys stink and drew the game :/")
			break
		}
		if turn == "X" {
			turn = "O"
		} else if turn == "O" {
			turn = "X"
		}
	}

}
func check_draw(board [3][3]string) bool {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board); col++ {
			if board[row][col] == "_" {
				return false
			}
		}
	}
	return true
}

func check_win(board [3][3]string) bool {

	for index := 0; index < len(board); index++ {
		// check horiz
		if turn := board[index][0]; turn != "_" && turn == board[index][1] && turn == board[index][2] {
			return true
		}
		// check vert
		if turn := board[0][index]; turn != "_" && turn == board[1][index] && turn == board[2][index] {
			return true
		}
	}
	//check diag
	if turn := board[0][0]; turn != "_" && turn == board[1][1] && turn == board[2][2] {
		return true
	}
	if turn := board[0][2]; turn != "_" && turn == board[1][1] && turn == board[2][0] {
		return true
	}

	return false
}

func player_turn(turn string, board *[3][3]string) bool {
	var col, row int
	fmt.Print("type col then row: ")
	fmt.Scan(&col, &row)

	if board[row][col] == "_" {
		board[row][col] = turn
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
