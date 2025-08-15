## Connect Four (CLI, Go)

A simple, fast, and emoji-friendly Connect Four game you can play in your terminal. Built with Go, featuring clean separation between CLI and game logic.

### Features

- Human vs. human local play in the terminal
- Columns selected via A–G keys; press Q to quit
- Clear board rendering with emoji pieces (🔴 for Red, 🟡 for Yellow)
- Move validation (no out-of-bounds or into full columns)
- Win detection (horizontal, vertical, and both diagonals)
- Auto-reset after a win so you can immediately play again

---

## Quick start

### Prerequisites

- Go 1.20+ installed
- make (optional, but recommended)

### Clone and run

```bash
# Clone
git clone https://github.com/neilsmahajan/connect-four.git
cd connect-four

# Run with Make (builds then runs)
make run

# Or build only
make build
./bin/connect-four

# Or run directly with Go
go run ./cmd/app/main.go
```

---

## How to play

When prompted:

- Type a letter A–G to drop your piece in that column.
- Uppercase or lowercase both work.
- Type Q at any time to quit.

Turn order alternates between Red and Yellow. The board is displayed after each move. When someone connects four in a row, the winner is announced and the board resets automatically for a new game.

Example prompt flow:

```
Connect Four CLI
Input A-G to drop a piece in that column, or Q to quit:
It's Red's turn.
```

Board rendering uses columns A–G with a 6x7 grid, like:

```
  A    B    C    D    E    F    G
+----+----+----+----+----+----+----+
|    |    |    |    |    |    |    |
+----+----+----+----+----+----+----+
|    |    |    |    |    |    |    |
+----+----+----+----+----+----+----+
|    |    |    |    |    |    |    |
+----+----+----+----+----+----+----+
|    |    |    |    |    |    |    |
+----+----+----+----+----+----+----+
|    |    |    |    |    |    |    |
+----+----+----+----+----+----+----+
| 🔴 | 🟡 |    |    |    |    |    |
+----+----+----+----+----+----+----+
```

---

## Makefile targets

- `make build` — Build the binary at `./bin/connect-four`
- `make run` — Build and run the game
- `make clean` — Remove the `./bin` directory

---

## Project structure

- `cmd/app/main.go` — App entrypoint that wires the CLI to the game board
- `internal/cli/cli.go` — CLI loop: prompts, input handling, and rendering
- `internal/board/board.go` — Core game logic: board state, turns, moves, win checks
- `Makefile` — Build, run, and clean helpers

---

## Implementation notes

- Input validation ensures only A–G and Q are accepted.
- Moves drop to the lowest available cell in the chosen column.
- Win detection checks horizontal, vertical, and both diagonal directions after each move.
- After a win, the board is reinitialized and play can continue immediately.

---

## Troubleshooting

- Emoji rendering: If your terminal font doesn’t support emoji, you might see empty boxes. Use a font with emoji support (e.g., Apple Color Emoji on macOS) or adjust your terminal settings.
- Terminal width: The board uses a fixed-width layout; ensure your terminal is wide enough to avoid wrapping.
- Go version: If you run into build issues, verify you’re on Go 1.20+.

---

## Contributing

Issues and pull requests are welcome. For larger changes, please open an issue first to discuss what you’d like to change.

---

## Author

Neil Mahajan  
Email: neilsmahajan@gmail.com  
Links: https://links.neilsmahajan.com/

---

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
