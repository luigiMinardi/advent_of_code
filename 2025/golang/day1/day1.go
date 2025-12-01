// in case we do a generic main package this will be the naming of this one
// package year2025day1
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	// dial are the numbers 0 through 99 in order
	// L or R which indicates whether the rotation should be to the
	// left (toward lower numbers) or to the right (toward higher numbers)
	// the dial is a circle, turning the dial left from 0 one click
	// makes it point at 99. Similarly, turning the dial right from 99 one click
	// makes it point at 0.
	// The dial starts by pointing at 50.
	// The actual password is the number of times the dial is left pointing
	// at 0 after any rotation in the sequence

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	dial := 50
	dial_zero := 0

	for scanner.Scan() {
		text := scanner.Text()
		direction := text[0]
		movements, err := strconv.Atoi(text[1:])
		if err != nil {
			fmt.Println(err)
		}
		// 76 -> rune for L
		// 82 -> rune for R
		for i := 0; i < movements; i++ {
			if direction == 76 {
				dial--
			} else {
				dial++
			}
			if dial == 100 {
				dial = 0
			} else if dial == -1 {
				dial = 99
			}
		}
		if dial == 0 {
			dial_zero++
		}
	}
	fmt.Println(dial_zero)
}

func part2() {
	// dial are the numbers 0 through 99 in order
	// L or R which indicates whether the rotation should be to the
	// left (toward lower numbers) or to the right (toward higher numbers)
	// the dial is a circle, turning the dial left from 0 one click
	// makes it point at 99. Similarly, turning the dial right from 99 one click
	// makes it point at 0.
	// The dial starts by pointing at 50.
	// The actual password is the number of times the dial is pointing
	// at 0 after any rotation in the sequence

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	dial := 50
	dial_zero := 0

	for scanner.Scan() {
		text := scanner.Text()
		direction := text[0]
		movements, err := strconv.Atoi(text[1:])
		if err != nil {
			fmt.Println(err)
		}
		// 76 -> rune for L
		// 82 -> rune for R
		for i := 0; i < movements; i++ {
			if direction == 76 {
				dial--
			} else {
				dial++
			}
			if dial == 100 {
				dial = 0
			} else if dial == -1 {
				dial = 99
			}
			if dial == 0 {
				dial_zero++
			}
		}
	}
	fmt.Println(dial_zero)
}
func main() {
	part2()
}
