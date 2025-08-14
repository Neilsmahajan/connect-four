package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/neilsmahajan/connect-four/internal/board"
)

func Run(board *board.Board) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Connect Four CLI")
		fmt.Println("Input A-G to drop a piece in that column, or Q to quit:")
		fmt.Printf("It's %s's turn.\n", getColorString(board.Turn))
		board.DrawBoard()
		for scanner.Scan() {

			input := scanner.Text()
			if input == "Q" || input == "q" {
				fmt.Println("Quitting the game.")
				return
			}
			if input >= "A" && input <= "G" || input >= "a" && input <= "g" {
				fmt.Printf("You selected column %s\n", input)
				// Here you would typically call a function to handle the game logic
			} else {
				fmt.Println("Invalid input. Please enter a column (A-G) or Q to quit.")
			}
			fmt.Println("Input A-G to drop a piece in that column, or Q to quit:")
			fmt.Printf("It's %s's turn.\n", getColorString(board.Turn))
			board.DrawBoard()
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
	}
}

func getColorString(color board.Color) string {
	switch color {
	case board.Red:
		return "Red"
	case board.Yellow:
		return "Yellow"
	case board.Empty:
		return "Empty"
	default:
		return "Unknown"
	}
}
