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

func part1(instructions []string) int {
	acc, _ := runProgram(instructions)
	return acc
}

func part2(instructions []string) int {
	acc := 0
	isCorrupted := false

	for i, instruction := range instructions {
		modifiedInstructions := make([]string, len(instructions))
		copy(modifiedInstructions, instructions)
		operation, _ := readInstruction(instruction)
		if operation == "nop" {
			modifiedInstructions[i] = strings.Replace(modifiedInstructions[i], "nop", "jmp", 1)
		} else if operation == "jmp" {
			modifiedInstructions[i] = strings.Replace(modifiedInstructions[i], "jmp", "nop", 1)
		}
		acc, isCorrupted = runProgram(modifiedInstructions)
		if !isCorrupted {
			break
		}
	}

	return acc
}

func runProgram(instructions []string) (acc int, isCorrupted bool) {
	readInstructions := make(map[int]int)
	i := 0

	for {
		if readInstructions[i] == 1 {
			isCorrupted = true
			break
		}
		if i >= len(instructions) {
			break
		}
		operation, argument := readInstruction(instructions[i])
		readInstructions[i] = 1
		switch operation {
		case "acc":
			acc += argument
			i++
		case "jmp":
			i += argument
		case "nop":
			i++
		}
	}

	return
}

func readInstruction(instruction string) (operation string, argument int) {
	split := strings.Split(instruction, " ")
	operation = split[0]
	argument, _ = strconv.Atoi(split[1])

	return
}
