package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := os.Args[1]
	number, err := strconv.Atoi(n)
	if err != nil {
		fmt.Printf("%s is not a valid number!\n", n)
		os.Exit(1)
	}

	l := []int{5, 7, 11, 48, 53, 63, 92, 102, 123, 124, 142, 161, 162, 167, 191, 228, 268, 274, 300, 306, 367, 410, 428, 438, 465, 466, 491, 512, 514, 549, 568, 593, 631, 633, 668, 670, 702, 718, 743, 756, 771, 780, 800, 802, 834, 875, 911, 926, 930, 939}
	fmt.Println(findNumber(number, l))
}

func findNumber(num int, list []int) bool {
	size := len(list)
	medPos := size / 2
	medValue := list[medPos]

	if medValue == num {
		return true
	}
	if size == 1 {
		return false
	}

	if medValue > num {
		newList := list[:medPos]
		return findNumber(num, newList)
	}

	newList := list[medPos:]
	return findNumber(num, newList)
}
