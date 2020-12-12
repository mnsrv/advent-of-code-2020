package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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
	input = append(input, 0)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	part1 := part1(input)
	part2 := part2(input)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func part1(input []int) (part1 int) {
	prev := 0
	one := 0
	two := 0
	three := 0

	for _, num := range input {
		if num-prev == 1 {
			one++
		} else if num-prev == 2 {
			two++
		} else if num-prev == 3 {
			three++
		}
		prev = num
	}
	part1 = one * three
	return
}

func part2(input []int) (part2 int) {
	part2 = 1
	prev := 0
	count := 0

	for _, num := range input {
		if num-prev == 3 {
			if count > 2 {
				combinations := 1
				if count == 3 {
					combinations = 2
				} else if count == 4 {
					combinations = 4
				} else if count == 5 {
					combinations = 7
				}
				part2 *= combinations
			}
			count = 1
		} else {
			count++
		}
		prev = num
	}

	return
}

// 2^(len-2)
// 2^(len-1)

// 0       4           2              0           0        0
// 0 1     4 5 6 7     10 11 12       15 16       19       22
// 0 1     4 5 6 7     10    12       15 16       19       22
// 0 1     4 5   7     10 11 12       15 16       19       22
// 0 1     4 5   7     10    12       15 16       19       22
// 0 1     4   6 7     10 11 12       15 16       19       22
// 0 1     4   6 7     10    12       15 16       19       22
// 0 1     4     7     10 11 12       15 16       19       22
// 0 1     4     7     10    12       15 16       19       22

// 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52

// 0 1 2 3 4     7 8 9 10 11       14       17 18 19 20       23 24 25       28       31 32 33 34 35       38 39       42       45 46 47 48 49       52

// 4             8                          4                 2                       8                                         8

// 0 1 2 3 4     7 8 9 10 11       14       17 18 19 20       23 24 25       28       31 32 33 34 35       38 39       42       45 46 47 48 49      (52)

// 0 1 2 3 4     7 8 9 10 11       14       17 18 19 20       23 24 25       28       31 32 33 34 35       38 39       42       45 46 47    49      (52)

// 0 1 2 3 4     7 8 9 10 11       14       17 18 19 20       23 24 25       28       31 32 33 34 35       38 39       42       45 46    48 49      (52)

// 0 1 2 3 4     7 8 9 10 11       14       17 18 19 20       23 24 25       28       31 32 33 34 35       38 39       42       45 46       49      (52)

// 0 1 2 3 4     7 8 9 10 11       14       17 18 19 20       23 24 25       28       31 32 33 34 35       38 39       42       45    47 48 49      (52)

// 0     3 4     7     10 11       14       17       20       23    25       28       31       34 35       38 39       42       45 46    48 49      (52)

// 0     3 4     7     10 11       14       17       20       23    25       28       31       34 35       38 39       42       45 46       49      (52)

// 0     3 4     7     10 11       14       17       20       23    25       28       31       34 35       38 39       42       45    47 48 49      (52)

// 0     3 4     7     10 11       14       17       20       23    25       28       31       34 35       38 39       42       45    47    49      (52)

// 0     3 4     7     10 11       14       17       20       23    25       28       31       34 35       38 39       42       45       48 49      (52)

// 1- 2

// 4 – 4
// 00
// 01
// 10
// 11

// 5 – 7
// -----000
// 001
// 010
// 011
// 100
// 101
// 110
// 111
