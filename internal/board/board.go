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
	result := "\n  A    B    C    D    E    F    G"
	for r := range board.Grid {
		result += "\n+----+----+----+----+----+----+----+\n"
		result += "|"
		for c := range board.Grid[r] {
			switch board.Grid[r][c].Color {
			case Red:
				result += " 🔴 |"
			case Yellow:
				result += " 🟡 |"
			case Empty:
				result += "    |"
			}
		}
	}
	result += "\n+----+----+----+----+----+----+----+\n"
	fmt.Println(result)
}

func (board *Board) CheckIfValidTurn(col int) bool {
	if col < 0 || col >= 7 {
		return false
	}
	if board.Grid[0][col].Color != Empty {
		return false
	}
	return true
}

func (board *Board) GetColorString(color Color) string {
	switch color {
	case Red:
		return "Red"
	case Yellow:
		return "Yellow"
	case Empty:
		return "Empty"
	default:
		return "Unknown"
	}
}

func switchTurn(current Color) Color {
	if current == Red {
		return Yellow
	}
	return Red
}

func (board *Board) CheckWin() (bool, Color) {
	// Check horizontal, vertical, and diagonal for a win
	for r := range board.Grid {
		for c := range board.Grid[r] {
			if board.Grid[r][c].Color == Empty {
				continue
			}
			color := board.Grid[r][c].Color
			// Check horizontal
			if c <= 3 && board.Grid[r][c+1].Color == color && board.Grid[r][c+2].Color == color && board.Grid[r][c+3].Color == color {
				return true, color
			}
			// Check vertical
			if r <= 2 && board.Grid[r+1][c].Color == color && board.Grid[r+2][c].Color == color && board.Grid[r+3][c].Color == color {
				return true, color
			}
			// Check diagonal (left to right)
			if r <= 2 && c <= 3 && board.Grid[r+1][c+1].Color == color && board.Grid[r+2][c+2].Color == color && board.Grid[r+3][c+3].Color == color {
				return true, color
			}
			// Check diagonal (right to left)
			if r <= 2 && c >= 3 && board.Grid[r+1][c-1].Color == color && board.Grid[r+2][c-2].Color == color && board.Grid[r+3][c-3].Color == color {
				return true, color
			}
		}
	}

	return false, Empty
}

func (board *Board) DropPiece(col int) error {
	if !board.CheckIfValidTurn(col) {
		return fmt.Errorf("invalid move")
	}
	for r := len(board.Grid) - 1; r >= 0; r-- {
		if board.Grid[r][col].Color == Empty {
			board.Grid[r][col].Color = board.Turn
			win, winner := board.CheckWin()
			if win {
				board.DrawBoard()
				fmt.Printf("%s wins!\n", board.GetColorString(winner))
				// Reset the board for a new game
				*board = *NewBoard()
				return nil
			}
			board.Turn = switchTurn(board.Turn)
			return nil
		}
	}
	return fmt.Errorf("column is full")
}
