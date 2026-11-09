package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// scene is the map 12h x 36w
const Height = 12 // line count
const Width = 36  // column count

// 12 tiles high column
type Column [Height]bool

type Scene struct {
	// width and height should be inverted
	Map      []Column
	BirdY    int
	Velocity int
	Score    int
}

func NewScene() Scene {
	return Scene{
		Map:      make([]Column, Width),
		BirdY:    8,
		Velocity: 0,
		Score:    0,
	}
}

func (state *Scene) Tick(jump bool) {
	state.Score++

	// shifts whole map to the left, and add a new column to the right
	state.Map = append(state.Map[1:], nextColumn(state))

	state.BirdY += state.Velocity

	if state.Velocity > -1 {
		state.Velocity--
	}

	if jump {
		state.Velocity = 2
	}
}

func nextColumn(state *Scene) Column {
	var col Column

	columnsAnchor := (state.Score + 11)

	if columnsAnchor%12 == 0 {
		// generates pipes
		openingHeight := rand.Int31n(7) + 3
		for i := 0; i < 12; i++ {
			// 5 tiles tall opening = < 3 tiles from the center of the opening
			isCloseToOpening := math.Abs(float64(i-int(openingHeight))) < 3
			if isCloseToOpening {
				col[i] = false
			} else {
				col[i] = true
			}
		}
	}

	return col
}

func main() {
	// local variable for state tracking
	scene := NewScene()

	fmt.Println("Flappy Bird")

	t := time.NewTicker(1 * time.Second)
	for range t.C {
		scene.Tick(false)
		render(scene)
	}
}
