package players

import (
	"fmt"
	"math"
	"math/rand"

	"github.com/jccroft1/tictactoebot/game"
)

type TQPlayer struct {
	Q             map[game.Board]map[game.Move]float32
	LearningRate  float32
	ValueDiscount float32
	QInit         float32
	DebugMode     bool
}

func (p *TQPlayer) Init() {
	p.Q = make(map[game.Board]map[game.Move]float32)
}

func (p *TQPlayer) InitBoard(b game.Board) {
	p.Q[b] = make(map[game.Move]float32)

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if b[r][c] != 0 {
				continue
			}

			p.Q[b][game.Move{r, c}] = p.QInit
		}
	}
}

func (p TQPlayer) Move(g game.Game) game.Move {
	if _, exists := p.Q[g.Board]; !exists {
		p.InitBoard(g.Board)
	}

	if p.DebugMode {
		fmt.Println(g.Board)
		fmt.Println(p.Q[g.Board])
	}

	bestMoves := make([]game.Move, 0)
	bestQ := float32(-1)
	for m, q := range p.Q[g.Board] {
		if q < bestQ {
			continue
		}

		if q > bestQ {
			bestMoves = []game.Move{m}
			bestQ = q
		}

		if q == bestQ {
			bestMoves = append(bestMoves, m)
		}
	}

	if len(bestMoves) == 0 {
		return bestMoves[0]
	}

	// pick a random best move
	return bestMoves[rand.Intn(len(bestMoves))]
}

// Learn takes a finished Game and adjusts Q values
func (p *TQPlayer) Learn(g game.Game) float64 {
	if g.Result == game.Unfinished {
		return 0
	}

	var qChange float64

	var q1, q2 float32
	if g.Result == game.Draw {
		q1, q2 = 0.5, 0.5
	} else {
		q1, q2 = 1, 0
	}

	lastMove := true
	for i := len(g.History) - 1; i >= 0; i = i - 2 {
		move := g.History[i]
		if lastMove {
			g.Board[move.Row][move.Column] = 0
			if _, exists := p.Q[g.Board]; !exists {
				p.InitBoard(g.Board)
			}
			qChange += math.Abs(float64(p.Q[g.Board][move] - q1))
			if p.DebugMode && math.Abs(float64(p.Q[g.Board][move]-q1)) != 0 {
				fmt.Println(g.Board)
				fmt.Println(move)
				fmt.Println("From: ", p.Q[g.Board][move], " To: ", q1)
			}
			p.Q[g.Board][move] = q1
			max := float32(-1)
			for _, q := range p.Q[g.Board] {
				if q > max {
					max = q
				}
			}
			q1 = max

			move = g.History[i-1]
			g.Board[move.Row][move.Column] = 0
			if _, exists := p.Q[g.Board]; !exists {
				p.InitBoard(g.Board)
			}
			qChange += math.Abs(float64(p.Q[g.Board][move] - q2))
			if p.DebugMode && math.Abs(float64(p.Q[g.Board][move]-q2)) != 0 {
				fmt.Println(g.Board)
				fmt.Println(move)
				fmt.Println("From: ", p.Q[g.Board][move], " To: ", q2)
			}
			p.Q[g.Board][move] = q2
			max = float32(-1)
			for _, q := range p.Q[g.Board] {
				if q > max {
					max = q
				}
			}
			q2 = max

			lastMove = false
		} else {
			g.Board[move.Row][move.Column] = 0
			if _, exists := p.Q[g.Board]; !exists {
				p.InitBoard(g.Board)
			}
			qChange += math.Abs(float64(p.Q[g.Board][move] - q1))
			if p.DebugMode && math.Abs(float64(p.Q[g.Board][move]-q1)) != 0 {
				fmt.Println(g.Board)
				fmt.Println(move)
				fmt.Println("From: ", p.Q[g.Board][move], " To: ", q1)
			}
			p.Q[g.Board][move] = (1-p.LearningRate)*p.Q[g.Board][move] + p.LearningRate*p.ValueDiscount*q1
			max := float32(-1)
			for _, q := range p.Q[g.Board] {
				if q > max {
					max = q
				}
			}
			q1 = max

			if i == 0 {
				continue
			}

			move := g.History[i-1]
			g.Board[move.Row][move.Column] = 0
			if _, exists := p.Q[g.Board]; !exists {
				p.InitBoard(g.Board)
			}
			qChange += math.Abs(float64(p.Q[g.Board][move] - q2))
			if p.DebugMode && math.Abs(float64(p.Q[g.Board][move]-q2)) != 0 {
				fmt.Println(g.Board)
				fmt.Println(move)
				fmt.Println("From: ", p.Q[g.Board][move], " To: ", q2)
			}
			p.Q[g.Board][move] = (1-p.LearningRate)*p.Q[g.Board][move] + p.LearningRate*p.ValueDiscount*q2

			max = float32(-1)
			for _, q := range p.Q[g.Board] {
				if q > max {
					max = q
				}
			}
			q2 = max
		}
	}

	return qChange / float64(len(g.History))
}
