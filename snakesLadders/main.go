package main

import "fmt"

func main() {
	var game SnakesLadders
	game.NewGame()
	fmt.Println(game.Play(1, 1))
	fmt.Println(game.Play(1, 5))
	fmt.Println(game.Play(6, 2))
	fmt.Println(game.Play(1, 1))
}

type GamePiece struct {
	Location int
	Name     string
}

func NewGamePiece(name string) GamePiece {
	return GamePiece{
		Location: 0,
		Name:     name,
	}
}

// SnakesLadders - Holds state for the game
type SnakesLadders struct {
	Turn          int
	Players       []GamePiece
	Board         map[int]int
	CurrentPlayer *GamePiece
}

// NewGame - Method that starts a new game for the SnakesLadders struct
func (sl *SnakesLadders) NewGame() {
	sl.Players = []GamePiece{
		NewGamePiece("Player 1"),
		NewGamePiece("Player 2"),
	}
	sl.Turn = 0
	sl.CurrentPlayer = &sl.Players[sl.Turn]
	sl.Board = map[int]int{2: 38, 7: 14, 8: 31, 15: 26, 16: 6, 21: 42, 28: 84, 36: 44, 46: 25, 49: 11, 51: 67, 62: 19, 64: 60, 71: 91, 74: 53, 78: 98, 87: 94, 89: 68, 92: 88, 95: 75, 99: 80}
}

// Play - Method that makes turn given a dice roll from die1 and die2 for the SnakesLadders struct
// Player that dice is for should automatically be determined based on rules
func (sl *SnakesLadders) Play(die1, die2 int) string {
	for i := range sl.Players {
		if sl.Players[i].Location >= 100 {
			return "Game over!"
		}
	}
	sl.CurrentPlayer.Location += die1 + die2
	if sl.CurrentPlayer.Location == 100 {
		return sl.CurrentPlayer.Name + " Wins!"
	}
	if sl.CurrentPlayer.Location > 100 {
		sl.CurrentPlayer.Location = 100 - (sl.CurrentPlayer.Location - 100)
	}
	if loc, ok := sl.Board[sl.CurrentPlayer.Location]; ok {
		sl.CurrentPlayer.Location = loc
	}

	r := fmt.Sprintf("%v is on square %v", sl.CurrentPlayer.Name, sl.CurrentPlayer.Location)
	if die1 != die2 {
		sl.Turn++
		if sl.Turn > len(sl.Players)-1 {
			sl.Turn = 0
		}
		sl.CurrentPlayer = &sl.Players[sl.Turn]
	}
	return r
}
