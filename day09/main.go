package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println(err)
	}
	inputString := strings.Split(string(data), "\n")
	input := make([]int, 0, len(inputString))
	for _, str := range inputString {
		n, _ := strconv.Atoi(str)
		input = append(input, n)
	}

	part1 := part1(input)
	part2 := part2(input, part1)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func part1(input []int) (part1 int) {
	preambleSize := 25
	previousSize := 25
	previous := make([]int, previousSize)
	for i, num := range input {
		if i < preambleSize {
			previous[i] = num
		} else {
			if !isValid(previous, num) {
				part1 = num
				break
			}
			previous[i%previousSize] = num
		}
	}
	return
}

func part2(input []int, part1 int) (part2 int) {
	l := 0
	sum := 0
	for i, num := range input {
		sum += num
		for sum > part1 {
			sum -= input[l]
			l++
		}
		if sum > part1 {
			sum -= input[l]
			l++
		}
		if sum == part1 {
			min := math.MaxInt64
			max := 0
			for _, n := range input[l : i+1] {
				if n < min {
					min = n
				}
				if n > max {
					max = n
				}
			}
			part2 = min + max
			break
		}
	}
	return
}

func isValid(previous []int, num int) bool {
	for i, num1 := range previous {
		for j, num2 := range previous {
			if i == j {
				continue
			}
			if num1+num2 == num {
				return true
			}
		}
	}
	return false
}
