package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {

	// var game = Game{
	// 	Rounds: []Round{
	// 		{
	// 			Opponent: Rock,
	// 			Me:       Paper,
	// 		},
	// 		{
	// 			Opponent: Paper,
	// 			Me:       Rock,
	// 		},
	// 		{
	// 			Opponent: Scissors,
	// 			Me:       Scissors,
	// 		},
	// 	},
	// }

	game := NewGameFromFileV2("./2022/02/input")

	game.CalculateScore()

	fmt.Println(game.Score)

}

type Game struct {
	Rounds []Round
	Score  int
}

func NewGameFromFile(file string) *Game {
	var game Game

	readData, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	splitData := bytes.Split(readData, []byte("\n"))

	for _, data := range splitData {
		var round Round

		if len(data) == 0 {
			continue
		}

		stringData := string(data)
		splitString := strings.Split(stringData, " ")

		round.Opponent = EncodeShape(splitString[0])
		round.Me = EncodeShape(splitString[1])

		game.Rounds = append(game.Rounds, round)
	}

	return &game
}

func NewGameFromFileV2(file string) *Game {
	var game Game

	readData, err := os.ReadFile(file)
	if err != nil {
		return nil
	}

	splitData := bytes.Split(readData, []byte("\n"))

	for _, data := range splitData {
		var round Round

		if len(data) == 0 {
			continue
		}

		stringData := string(data)
		splitString := strings.Split(stringData, " ")

		round.Opponent = EncodeShape(splitString[0])
		round.Me = ResultToShape(splitString[1], round.Opponent)

		game.Rounds = append(game.Rounds, round)
	}

	return &game
}

func (g *Game) CalculateScore() {
	var score int

	for _, r := range g.Rounds {

		r.CalculateResult()

		score = score + r.Me.Value() + r.Result.Value()
	}

	g.Score = score
}

type Round struct {
	Opponent Shape
	Me       Shape
	Result   Result
}

func (r *Round) CalculateResult() {

	switch r.Me {
	case Rock:
		if r.Opponent == Rock {
			r.Result = Draw
		} else if r.Opponent == Paper {
			r.Result = Lose
		} else if r.Opponent == Scissors {
			r.Result = Win
		}
	case Paper:
		if r.Opponent == Rock {
			r.Result = Win
		} else if r.Opponent == Paper {
			r.Result = Draw
		} else if r.Opponent == Scissors {
			r.Result = Lose
		}
	case Scissors:
		if r.Opponent == Rock {
			r.Result = Lose
		} else if r.Opponent == Paper {
			r.Result = Win
		} else if r.Opponent == Scissors {
			r.Result = Draw
		}
	}
}

type Result int

const (
	Win Result = iota
	Lose
	Draw
)

func (r Result) Value() int {
	switch r {
	case Win:
		return 6
	case Draw:
		return 3
	case Lose:
		return 0
	default:
		return 0
	}
}

func EncodeResult(in string) Result {
	switch in {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	default:
		return Win
	}
}

func EncodeShape(in string) Shape {
	switch in {
	case "A":
		return Rock
	case "X":
		return Rock
	case "B":
		return Paper
	case "Y":
		return Paper
	case "C":
		return Scissors
	case "Z":
		return Scissors
	default:
		return Rock // TODO: Should be an invalid type
	}
}

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

func (s Shape) Value() int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}

	return 0
}

func ResultToShape(resultString string, opponentShape Shape) Shape {
	var desiredResult = EncodeResult(resultString)

	switch desiredResult {
	case Win:
		if opponentShape == Rock {
			return Paper
		} else if opponentShape == Paper {
			return Scissors
		} else if opponentShape == Scissors {
			return Rock
		}
	case Draw:
		return opponentShape
	case Lose:
		if opponentShape == Rock {
			return Scissors
		} else if opponentShape == Paper {
			return Rock
		} else if opponentShape == Scissors {
			return Paper
		}
	}

	return Rock // TODO: Make this an invalid type
}
