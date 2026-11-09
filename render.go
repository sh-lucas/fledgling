package main

import "fmt"

func render(state Scene) {
	// limpa 13 linhas acima
	fmt.Printf("\033[%dA\033[J", 15)

	fmt.Printf("Score: %04d | ==========================================================\n", state.Score)

	for line := 11; line >= 0; line-- {
		for column := 0; column < 36; column++ {
			// bird =) always on second column
			if line == state.BirdY && column == 1 {
				fmt.Printf("🐦")
				continue
			}

			pos := state.Map[column][line]
			if pos {
				fmt.Printf("██")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("Score:", state.Score, "Speed:", state.Velocity)
}
