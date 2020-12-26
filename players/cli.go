package players

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jccroft1/tictactoebot/game"
)

type CliPlayer struct {
	Name string
}

func (p CliPlayer) Move(b game.Game) game.Move {
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("%v's move\n", p.Name)

	fmt.Println("Board: ")
	fmt.Println(b.Board)

	var row, col int
	var rowText, colText string
	for {
		err := errors.New("first time")
		for err != nil {
			fmt.Print("Enter row: ")
			rowText, err = reader.ReadString('\n')
			row, err = strconv.Atoi(strings.TrimSpace(rowText))
		}

		err = errors.New("first time")
		for err != nil {
			fmt.Print("Enter col: ")
			colText, err = reader.ReadString('\n')
			col, err = strconv.Atoi(strings.TrimSpace(colText))
		}

		if 0 < row && row <= 3 && 0 < col && col <= 3 && b.Board[row-1][col-1] == 0 {
			break
		}
		fmt.Println("invalid move")
	}

	return game.Move{row - 1, col - 1}
}
