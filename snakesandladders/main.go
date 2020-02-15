package kata

import (
	"fmt"
)

var snakes = map[int]int {
	16: 6,
	46: 25,
	49: 11,
	62: 19,
	64: 60,
	74: 53,
	89: 68,
	92: 88,
	95: 75,
	99: 80,
}

var ladders = map[int]int {
	2: 38,
	7: 14,
	8: 31,
	15: 26,
	21: 42,
	28: 84,
	36: 44,
	51: 67,
	71: 91,
	78: 98,
	87: 94,
}

type SnakesLadders struct {
	nextPlayer int
	players [2]int
}

func (sl *SnakesLadders) NewGame() {
	sl.nextPlayer = 0
	sl.players = [2]int{}
}

func (sl *SnakesLadders) Play(die1, die2 int) (ret string) {
	sum := die1 + die2
	curPlayer := sl.nextPlayer

	if sl.players[0] == 100 || sl.players[1] == 100 {
		return "Game over!"
	}

	if (die1 != die2) {
		if sl.nextPlayer == 1 {
			sl.nextPlayer = 0
		} else {
			sl.nextPlayer = 1
		}
	}

	sl.players[curPlayer] += sum

	if sl.players[curPlayer] == 100 {
		return fmt.Sprintf("Player %d Wins!", curPlayer + 1)
	}

	if sl.players[curPlayer] > 100 {
		sl.players[curPlayer] = 100 - (sl.players[curPlayer] - 100)
	}

	if snake, ok := snakes[sl.players[curPlayer]]; ok {
		sl.players[curPlayer] = snake
	}

	if ladder, ok := ladders[sl.players[curPlayer]]; ok {
		sl.players[curPlayer] = ladder
	}

	return fmt.Sprintf("Player %d is on square %d", curPlayer + 1, sl.players[curPlayer])
}
