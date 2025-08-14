package board

import "fmt"

type Color int

const (
	Red Color = iota
	Yellow
	Empty
)

type Piece struct {
	Color Color
}

type Position struct {
	Row, Col int
}

type Board struct {
	Grid [6][7]Piece
	Turn Color
}

func NewBoard() *Board {
	board := &Board{Turn: Red}
	for r := range board.Grid {
		for c := range board.Grid[r] {
			board.Grid[r][c] = Piece{Color: Empty}
		}
	}
	return board
}

func (board *Board) DrawBoard() {
	result := "\n  A   B   C   D   E   F   G"
	for r := range board.Grid {
		result += "\n+---+---+---+---+---+---+---+\n"
		result += "|"
		for c := range board.Grid[r] {
			switch board.Grid[r][c].Color {
			case Red:
				result += " R |"
			case Yellow:
				result += " Y |"
			case Empty:
				result += "   |"
			}
		}
	}
	result += "\n+---+---+---+---+---+---+---+\n"
	fmt.Println(result)
}
