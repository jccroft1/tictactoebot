package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jccroft1/tictactoebot/game"
	"github.com/jccroft1/tictactoebot/players"
)

const testingIterations = 50000

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
	r1 := players.RandomPlayer{}
	tq1 := players.TQPlayer{
		LearningRate:  0.9,
		ValueDiscount: 0.95,
		QInit:         0.75,
	}
	tq1.Init()

	var g game.Game
	// Train
	for i := 0; i < testingIterations; i++ {
		// g = game.New(tq1, r1)
		// tq1.Learn(g)

		g = game.New(tq1, tq1)
		tq1.Learn(g)
		// qChangePerGame := tq1.Learn(g)
		// fmt.Println(qChangePerGame)

		if i%(testingIterations/10) == 0 {
			fmt.Println("\nIteration: ", i/(testingIterations/10))
			test(tq1, mm1)
			test(tq1, r1)
		}
	}

	// g = game.New(tq1, tq1)
	// tq1.DebugMode = true
	// tq1.Learn(g)
	// tq1.DebugMode = false
	// g = game.New(tq1, tq1)
	// tq1.DebugMode = true
	// tq1.Learn(g)

	g = game.New(tq1, p1)

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

func test(subject, reference game.Player) {
	var draws, wins, losses int

	for i := 0; i < 50; i++ {
		g := game.New(subject, reference)
		switch g.Result {
		case game.Draw:
			draws++
		case game.OWin:
			wins++
		case game.XWin:
			losses++
		}

		g = game.New(reference, subject)
		switch g.Result {
		case game.Draw:
			draws++
		case game.OWin:
			losses++
		case game.XWin:
			wins++
		}
	}

	fmt.Println("Draws: ", draws)
	fmt.Println("Wins: ", wins)
	fmt.Println("Losses: ", losses)
}
