package main

import (
	"github.com/neilsmahajan/connect-four/internal/board"
	"github.com/neilsmahajan/connect-four/internal/cli"
)

func main() {
	board := board.NewBoard()
	println(board.DrawBoard())
	cli.Run()
}
