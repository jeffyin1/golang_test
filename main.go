package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	test()
}

const DEFAULT_RANK Rank = 1

var startScore int = 100

type Rank int

type Game struct {
	Score int
	Rank  Rank
}

func (g *Game) addScore(a int) int {
	g.Score += a
	return g.Score
}

func (r *Rank) levelUp() {
	*r = *r + 1
}

func test() {
	game := Game{Score: startScore, Rank: DEFAULT_RANK}
	game.addScore(10)
	fmt.Println(game.Score)
	game.Rank.levelUp()
	fmt.Println(game.Rank)
}
