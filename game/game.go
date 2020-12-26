package game

import (
	"errors"
	"fmt"
)

// Rules
// o players first

type Result int16

const (
	Unfinished Result = iota
	OWin
	XWin
	Draw
)

func (s Result) String() string {
	switch s {
	case Unfinished:
		return "Unfinished"
	case OWin:
		return "O Wins"
	case XWin:
		return "X Wins"
	case Draw:
		return "Draw"
	}

	return "?"
}

// Square
// blank = 0
// o = 1
// x = 2
type Square uint8

const (
	Empty Square = iota
	O
	X
)

func (s Square) String() string {
	switch s {
	case Empty:
		return " "
	case O:
		return "O"
	case X:
		return "X"
	}

	return "?"
}

// Board is a 3x3 nested array
// indexed by row, then column
type Board [3][3]Square

func (b Board) String() string {
	return fmt.Sprintf("%v|%v|%v\n%v|%v|%v\n%v|%v|%v\n",
		b[0][0], b[0][1], b[0][2],
		b[1][0], b[1][1], b[1][2],
		b[2][0], b[2][1], b[2][2])
}

// Move int represents position
// x or o is obvious from board
type Move struct {
	Row, Column int
}

func (m Move) IsValid(b Board) bool {
	return b[m.Row][m.Column] == 0
}

type Game struct {
	Board  Board
	Turn   int
	Result Result
}

func (g Game) CheckState(player Square, m Move) Result {
	if g.Turn == 9 {
		return Draw
	}

	win := false
	if g.Board[1][1] == player {
		if g.Board[0][0] == player && g.Board[2][2] == player {
			win = true
		}

		if g.Board[2][0] == player && g.Board[0][2] == player {
			win = true
		}
	}

	if g.Board[m.Row][0] == player && g.Board[m.Row][1] == player && g.Board[m.Row][2] == player {
		win = true
	}

	if g.Board[0][m.Column] == player && g.Board[1][m.Column] == player && g.Board[2][m.Column] == player {
		win = true
	}

	if win {
		if player == 1 {
			return OWin
		}

		return XWin
	}

	return Unfinished
}

func (g Game) Apply(m Move) (Game, error) {
	// check it's a valid move
	if !m.IsValid(g.Board) {
		return g, errors.New("invalid move")
	}

	// determine current player
	player := Square((g.Turn % 2) + 1)

	// update board
	g.Board[m.Row][m.Column] = player

	// increment turn counter
	g.Turn++

	// check for winner
	g.Result = g.CheckState(player, m)

	// return changed game
	return g, nil
}

type Player interface {
	Move(Game) Move
}

func New(p1, p2 Player) Game {
	g := Game{
		Board: Board{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}

	for i := 0; i < 9; i++ {
		move1 := p1.Move(g)
		g, _ = g.Apply(move1)
		if g.Result != Unfinished {
			break
		}

		move2 := p2.Move(g)
		g, _ = g.Apply(move2)

		if g.Result != Unfinished {
			break
		}
	}

	return g
}
