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

	for _, str := range input {
		first, second, letter, password := getData(str)
		part1 += checkValidityPart1(first, second, letter, password)
	}
	return part1
}

func part2(input []string) int {
	part1 := 0

	for _, str := range input {
		first, second, letter, password := getData(str)
		part1 += checkValidityPart2(first, second, letter, password)
	}
	return part1
}

func getData(str string) (first int, second int, letter string, password string) {
	strSplited := strings.Split(str, " ")
	times := strSplited[0]
	letter = strings.TrimRight(strSplited[1], ":")
	password = strSplited[2]
	timesSplited := strings.Split(times, "-")
	first, _ = strconv.Atoi(timesSplited[0])
	second, _ = strconv.Atoi(timesSplited[1])

	return
}

func checkValidityPart1(min int, max int, letter string, password string) int {
	isValid := 0

	count := strings.Count(password, letter)
	if count >= min && count <= max {
		isValid = 1
	}
	return isValid
}

func checkValidityPart2(min int, max int, letter string, password string) int {
	isValid := 0

	if (string(password[min-1]) == letter) != (string(password[max-1]) == letter) {
		isValid = 1
	}
	return isValid
}
