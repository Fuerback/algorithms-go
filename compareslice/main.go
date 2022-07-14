package main

import "fmt"

func main() {
	var names = []string{"F5", "F7", "C6", "H5", "G5"}
	var board = []string{"B4", "D4", "G5", "F5", "F7", "C6"}

	var boardMap = make(map[string]bool)

	for _, value := range board {
		boardMap[value] = true
	}

	var result []string

	for _, value := range names {
		if boardMap[value] {
			result = append(result, value)
		}
	}

	fmt.Println(result)
}
