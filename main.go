package main

import (
	"fmt"
)

func main() {
	board := [3][3]string{} // initialize the board

	player := "X"
	winner := ""

	// loop until there is a winner or the board is full
	for i := 0; i < 9 && winner == ""; i++ {
		fmt.Printf("Player %s's turn\n", player)

		// ask the player to input row and column
		var row, col int
		fmt.Print("Enter row (0-2): ")
		fmt.Scanln(&row)
		fmt.Print("Enter column (0-2): ")
		fmt.Scanln(&col)

		// check if the input is valid
		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid input, try again.")
			i--
			continue
		}

		// check if the cell is already occupied
		if board[row][col] != "" {
			fmt.Println("Cell already occupied, try again.")
			i--
			continue
		}

		// mark the cell with the player's symbol
		board[row][col] = player

		// check if the player wins
		if board[row][0] == player && board[row][1] == player && board[row][2] == player ||
			board[0][col] == player && board[1][col] == player && board[2][col] == player ||
			board[0][0] == player && board[1][1] == player && board[2][2] == player ||
			board[0][2] == player && board[1][1] == player && board[2][0] == player {
			winner = player
		}

		// switch to the other player
		if player == "X" {
			player = "O"
		} else {
			player = "X"
		}

		// print the board
		fmt.Println("*---------*")
		for row := 0; row < 3; row++ {
			fmt.Print("| ")
			for col := 0; col < 3; col++ {
				fmt.Printf("%s | ", board[row][col])
			}
			fmt.Println()
			fmt.Println("*---------*")
		}
	}

	// print the winner or tie
	if winner != "" {
		fmt.Printf("Player %s wins!\n", winner)
	}
}
