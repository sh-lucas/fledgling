package main

import "fmt"

func render(state Scene) {
	fmt.Println("========================================================================")
	for line := range Height {
		for column := range Width {
			if state.Map[line][column] {
				fmt.Printf("██")
			} else {
				fmt.Printf(" .")
			}
		}
		fmt.Printf("\n")
	}
}
