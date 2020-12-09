package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	input := strings.Split(string(data), "\n")
	bags := getBags(input)

	part1 := part1(bags)
	part2 := part2(bags)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func getBags(input []string) map[string][][]string {
	bags := make(map[string][][]string)

	for _, rule := range input {
		regexpColor, _ := regexp.Compile("(.+?) bag")
		color := regexpColor.FindStringSubmatch(rule)[1]
		regexpContains, _ := regexp.Compile("(\\d+) (.+?) bag")
		containsMatch := regexpContains.FindAllStringSubmatch(rule, -1)
		contains := make([][]string, len(containsMatch))
		for i, contain := range containsMatch {
			contains[i] = contain[1:]
		}
		bags[color] = contains
	}

	return bags
}

func hasShiny(bags map[string][][]string, color string) bool {
	if color == "shiny gold" {
		return true
	}
	for _, value := range bags[color] {
		if hasShiny(bags, value[1]) {
			return true
		}
	}
	return false
}

func part1(bags map[string][][]string) (part1 int) {
	for color := range bags {
		if hasShiny(bags, color) {
			part1++
		}
	}
	part1--
	return
}

func count(bags map[string][][]string, color string) int {
	sum := 1

	for _, contain := range bags[color] {
		num, _ := strconv.Atoi(contain[0])
		c := count(bags, contain[1])
		sum += num * c
	}

	return sum
}

func part2(bags map[string][][]string) (part2 int) {
	return count(bags, "shiny gold") - 1
}
