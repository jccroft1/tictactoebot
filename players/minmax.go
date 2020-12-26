package players

import (
	"fmt"

	"github.com/jccroft1/tictactoebot/game"
)

type EvaluatedMove struct {
	Move  game.Move
	Score int
}

type MinMaxPlayer struct {
	MaxCache map[game.Board]EvaluatedMove
	MinCache map[game.Board]EvaluatedMove
}

func (p *MinMaxPlayer) Init() {
	p.MaxCache = make(map[game.Board]EvaluatedMove)
	p.MinCache = make(map[game.Board]EvaluatedMove)
}

func (p MinMaxPlayer) Move(b game.Game) game.Move {
	ev := p.max(b)
	return ev.Move
}

func (p MinMaxPlayer) max(g game.Game) EvaluatedMove {
	if ev, exists := p.MaxCache[g.Board]; exists {
		return ev
	}

	validMoves := make(map[game.Move]EvaluatedMove)
	player := game.Square((g.Turn % 2) + 1)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			m := game.Move{r, c}

			if !m.IsValid(g.Board) {
				continue
			}

			ev := EvaluatedMove{
				Move: m,
			}

			g1, err := g.Apply(m)
			if err != nil {
				fmt.Println("tried to apply invalid move")
				continue
			}

			// Immdetialy win, obviously the best choice
			if (g1.Result == game.OWin && player == 1) || (g1.Result == game.XWin && player == 2) {
				ev.Score = 1
				p.MaxCache[g.Board] = ev
				return ev
			}

			if g1.Result == game.Draw {
				ev.Score = 0
			} else if g1.Result == game.Unfinished {
				// game is unfinished
				ev1 := p.min(g1)
				ev.Score = ev1.Score
			} else {
				// lost
				ev.Score = -1
			}

			validMoves[m] = ev
		}
	}

	currentMove := EvaluatedMove{
		Score: -2,
	}
	for _, ev := range validMoves {
		if ev.Score > currentMove.Score {
			currentMove = ev
		}
	}

	p.MaxCache[g.Board] = currentMove

	return currentMove
}

func (p MinMaxPlayer) min(g game.Game) EvaluatedMove {
	if ev, exists := p.MinCache[g.Board]; exists {
		return ev
	}

	validMoves := make(map[game.Move]EvaluatedMove)
	player := game.Square((g.Turn % 2) + 1)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			m := game.Move{r, c}

			if !m.IsValid(g.Board) {
				continue
			}

			ev := EvaluatedMove{
				Move: m,
			}

			g1, err := g.Apply(m)
			if err != nil {
				fmt.Println("tried to apply invalid move")
				continue
			}

			if (g1.Result == game.OWin && player == 1) || (g1.Result == game.XWin && player == 2) {
				ev.Score = -1
			} else if g1.Result == game.Draw {
				ev.Score = 0
			} else if g1.Result == game.Unfinished {
				// game is unfinished
				ev1 := p.max(g1)
				ev.Score = ev1.Score
			} else {
				// lost
				ev.Score = 1
			}

			validMoves[m] = ev
		}
	}

	currentMove := EvaluatedMove{
		Score: 2,
	}
	for _, ev := range validMoves {
		if ev.Score < currentMove.Score {
			currentMove = ev
		}
	}

	p.MinCache[g.Board] = currentMove

	return currentMove
}
