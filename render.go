package main

import "fmt"

func render(state Scene) {
	// limpa 13 linhas acima
	fmt.Printf("\033[%dA\033[J", 14)

	fmt.Printf("Score: %04d ============================================================\n", state.Score)

	for line := 11; line >= 0; line-- {
		for column := 0; column < 36; column++ {
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
}
