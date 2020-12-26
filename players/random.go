package players

import (
	"math/rand"

	"github.com/jccroft1/tictactoebot/game"
)

type RandomPlayer struct{}

func (p RandomPlayer) Move(b game.Game) game.Move {
	var validMoves []game.Move
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b.Board[r][c] == 0 {
				validMoves = append(validMoves, game.Move{r, c})
			}
		}
	}

	chosenMove := validMoves[rand.Intn(len(validMoves))]

	return chosenMove
}
