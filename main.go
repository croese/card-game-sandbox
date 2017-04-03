package main

import "fmt"

type Player interface {
}

type Game struct {
	players []*Player
}

func (g *Game) AddPlayer(p *Player) {
	g.players = append(g.players, p)
}

func main() {
	fmt.Println("Card Game Sandbox!")
}
