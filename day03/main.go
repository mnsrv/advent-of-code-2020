package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "\n")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func part1(input []string) int {
	return countTrees(input, 3, 1)
}

func part2(input []string) int {
	return countTrees(input, 1, 1) * countTrees(input, 3, 1) * countTrees(input, 5, 1) * countTrees(input, 7, 1) * countTrees(input, 1, 2)
}

func countTrees(input []string, right int, down int) int {
	trees := 0
	coords := make(map[point]int)
	maxX := len(input[0])
	maxY := len(input)

	for y, row := range input {
		for x, el := range row {
			if string(el) == "#" {
				coords[point{x, y}] = 1
			} else {
				coords[point{x, y}] = 0
			}
		}
	}

	x := 0
	y := 0
	for i := 0; i <= maxY; i++ {
		x = i * right
		y = i * down
		trees += coords[point{x % maxX, y}]
	}

	return trees
}
