package main

import (
	"fmt"
	"time"
)

type Game struct {
	Players map[IDType]*Player
	Objects map[IDType]*Object
	GameOver bool
}

// game := Game{
// 	Players: [],
// 	Objects: []
// }

func (g *Game) AddPlayer (id IDType) {
	// game.Players[id] = new Player()
}

func (g *Game) DelPlayer (id IDType) {
	// game.Players[id] = nil
}

func (g *Game) GameTicker (id IDType) {
	for !g.GameOver {
		// game.HandleEvent()
		for _, player := range g.Players {
			player.Update()
		}

		time.Sleep(MS_PER_TICK)
	}
}

func (g *Game) Start () {
	g.Players = make(map[IDType]*Player)
	g.Objects = make(map[IDType]*Object)

	g.GameOver = false
	fmt.Println("Game started on server!")

}

func (g *Game) HandleEvent () {

}