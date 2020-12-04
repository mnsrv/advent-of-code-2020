package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

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
	part1 := 0

Loop:
	for _, str1 := range input {
		for _, str2 := range input {
			expense1, _ := strconv.Atoi(str1)
			expense2, _ := strconv.Atoi(str2)

			sum := expense1 + expense2
			if sum == 2020 {
				part1 = expense1 * expense2
				break Loop
			}
		}
	}
	return part1
}

func part2(input []string) int {
	part2 := 0

Loop:
	for _, str1 := range input {
		for _, str2 := range input {
			for _, str3 := range input {
				expense1, _ := strconv.Atoi(str1)
				expense2, _ := strconv.Atoi(str2)
				expense3, _ := strconv.Atoi(str3)

				sum := expense1 + expense2 + expense3
				if sum == 2020 {
					part2 = expense1 * expense2 * expense3
					break Loop
				}
			}
		}
	}
	return part2
}
