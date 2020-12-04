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
	input := strings.Split(string(data), "\n\n")

	part1 := part1(input)
	part2 := part2(input)

	fmt.Println("part one:", part1)
	fmt.Println("part two:", part2)
}

func part1(input []string) int {
	part1 := 0

	for _, str := range input {
		passport := getData(str)
		part1 += checkFieldsPresent(passport)
	}

	return part1
}

func part2(input []string) int {
	part2 := 0

	for _, str := range input {
		passport := getData(str)
		part2 += (checkFieldsPresent(passport) * checkFieldsValid(passport))
	}

	return part2
}

func getData(str string) map[string]string {
	passport := make(map[string]string)

	fields := strings.Fields(str)
	for _, field := range fields {
		f := strings.Split(field, ":")
		passport[f[0]] = f[1]
	}

	return passport
}

func checkFieldsPresent(passport map[string]string) (isValid int) {
	isValid = 1
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	for _, field := range requiredFields {
		if _, ok := passport[field]; !ok {
			isValid = 0
			break
		}
	}

	return
}

func checkFieldsValid(passport map[string]string) (isValid int) {
	isValid = 1
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	ecl := map[string]int{"amb": 0, "blu": 0, "brn": 0, "gry": 0, "grn": 0, "hzl": 0, "oth": 0}

Loop:
	for _, field := range requiredFields {
		if value, ok := passport[field]; ok {
			switch field {
			case "byr":
				year, _ := strconv.Atoi(value)
				if year < 1920 || year > 2002 {
					isValid = 0
					break Loop
				}
			case "iyr":
				year, _ := strconv.Atoi(value)
				if year < 2010 || year > 2020 {
					isValid = 0
					break Loop
				}
			case "eyr":
				year, _ := strconv.Atoi(value)
				if year < 2020 || year > 2030 {
					isValid = 0
					break Loop
				}
			case "hgt":
				measure := string(value[len(value)-2:])
				height, _ := strconv.Atoi(value[:len(value)-2])
				if measure == "cm" {
					if height < 150 || height > 193 {
						isValid = 0
						break Loop
					}
				} else if measure == "in" {
					if height < 59 || height > 76 {
						isValid = 0
						break Loop
					}
				} else {
					isValid = 0
					break Loop
				}
			case "hcl":
				r, _ := regexp.Compile("#[0-9a-f]{6}")
				if !r.MatchString(value) {
					isValid = 0
					break Loop
				}
			case "ecl":
				if _, ok := ecl[value]; !ok {
					isValid = 0
					break Loop
				}
			case "pid":
				_, err := strconv.Atoi(value)
				if len(value) != 9 || err != nil {
					isValid = 0
					break Loop
				}

			}
		}
	}

	return
}
