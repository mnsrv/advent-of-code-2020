package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "\n\n")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func part1(input []string) int {
	part1 := 0

	for _, group := range input {
		m := make(map[rune]int)
		persons := strings.Split(group, "\n")
		for _, person := range persons {
			for _, answer := range person {
				m[answer] = 1
			}
		}
		part1 += len(m)
	}

	return part1
}

func part2(input []string) int {
	part2 := 0

	for _, group := range input {
		m := make(map[rune]int)
		persons := strings.Split(group, "\n")
		for _, person := range persons {
			for _, answer := range person {
				m[answer]++
				if m[answer] == len(persons) {
					part2++
				}
			}
		}
	}

	return part2
}
