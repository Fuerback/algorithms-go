package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))
	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s is not a valid number!\n", n)
			os.Exit(1)
		}
		numbers[i] = number
	}
	fmt.Println(quicksort(numbers))
}

func quicksort(numbers []int) []int {
	// verify if the list has more than one element
	if len(numbers) <= 1 {
		return numbers
	}

	// create a new list and copy the original
	n := make([]int, len(numbers))
	copy(n, numbers)

	// select an element to be the reference
	refIndex := len(n) / 2
	ref := n[refIndex]

	// now we need to remove this ref element from copy list, so we'll use the append func to do that
	// we have to copy all elements before refIndex and after refIndex, so that way we are avoiding the refElement
	n = append(n[:refIndex], n[refIndex+1:]...)

	// with copy list and refElement, we are going to partition in two lists, one with greatest numbers than reference and other with smallest numbers
	smallest, greatest := partition(n, ref)
	
	// ordering the lists recursively and return
	return append(
		append(quicksort(smallest), ref),
		quicksort(greatest)...)
}

func partition(n []int, ref int) (smallest []int, greatest []int) {
	for _, n := range n {
		if n <= ref {
			smallest = append(smallest, n)
		} else {
			greatest = append(greatest, n)
		}
	}
	return smallest, greatest
}
