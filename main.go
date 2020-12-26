package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jccroft1/tictactoebot/game"
	"github.com/jccroft1/tictactoebot/players"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	p1 := players.CliPlayer{
		Name: "Jack",
	}
	mm1 := players.MinMaxPlayer{}
	mm1.Init()
	// p2 := CliPlayer{
	// 	Name: "Jack",
	// }
	// r1 := players.RandomPlayer{}

	g := game.New(p1, mm1)
	fmt.Println(g.Board)
	fmt.Println(g.Result)

	// var o, x, draw int

	// for i := 0; i < 100000; i++ {
	// 	g := game.New(r1, mm1)
	// 	switch g.Result {
	// 	case game.Draw:
	// 		draw++
	// 	case game.OWin:
	// 		o++
	// 	case game.XWin:
	// 		x++
	// 	}
	// }

	// fmt.Println("Draws: ", draw)
	// fmt.Println("O Wins: ", o)
	// fmt.Println("X Wins: ", x)
}
