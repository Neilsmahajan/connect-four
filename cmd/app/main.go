package main

import "github.com/neilsmahajan/connect-four/internal/board"

func main() {
	board := board.NewBoard()
	println(board.DrawBoard())
}
