package cli

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Connect Four CLI")
		fmt.Println("Input A-G to drop a piece in that column, or Q to quit:")
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
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "Error reading input:", err)
			return
		}
	}
}
