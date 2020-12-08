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
	input := strings.Split(string(data), "\n")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func part1(input []string) int {
	part1 := 0

	for _, str := range input {
		seatID := getSeatID(str)
		if seatID > part1 {
			part1 = seatID
		}
	}

	return part1
}

func getSeatID(pass string) int {
	passReplacer := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	pass2 := passReplacer.Replace(pass)
	pass10, _ := strconv.ParseInt(pass2, 2, 64)

	return int(pass10)
}

func part2(input []string) int {
	part2 := 0
	m := make(map[int]int, 1024)
	for i := 0; i < 1023; i++ {
		m[i] = 1
	}

	for _, str := range input {
		seatID := getSeatID(str)
		// the only missing seatID is mine
		delete(m, seatID)
	}

	keys := make([]int, 0)
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		// seatID +1 and -1 exists
		if m[k-1] != 1 && m[k+1] != 1 {
			part2 = k
		}
	}

	return part2
}
