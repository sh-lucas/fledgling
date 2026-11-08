package main

import (
	"fmt"
	"time"
)

// scene is the map 12h x 36w
const Height = 12
const Width = 36

type Scene struct {
	Map   [Height][Width]bool
	BirdY int
	// builds up every tick towards -1; adds to BirdY
	Velocity int
	// y distance from start and how many seconds have passed
	Score int
}

func (state *Scene) Tick() {
	state.Score++

	if state.Score%2 == 0 {
		state.BirdY--
	}

	if state.Score%12 == 0 {
		spawnPipes(state)
	}
}

func spawnPipes(state *Scene) {

}

func main() {
	// local variable for state tracking
	var scene Scene
	fmt.Println("Flappy Bird")

	t := time.NewTicker(1 * time.Second)
	for range t.C {
		render(scene)
	}
}
