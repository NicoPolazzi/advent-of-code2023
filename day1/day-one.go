package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var sumPartOne, sumPartTwo int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		sumPartOne += partOne(line)
		sumPartTwo += partTwo(line)
	}

	fmt.Println("The sum from the first part is:", sumPartOne)
	fmt.Println("The sum from the first part is:", sumPartTwo)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func partOne(line string) int {
	var values []string

	for _, r := range line {
		if unicode.IsDigit(r) {
			values = append(values, string(r))
		}
	}

	calibrationValue, _ := strconv.Atoi(values[0] + values[len(values)-1])
	return calibrationValue
}

func partTwo(line string) int {
	digits := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"0":     "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
	}

	var values []string

	for i := range line {
		for key, value := range digits {
			if strings.HasPrefix(line[i:], key) {
				values = append(values, value)
			}
		}
	}

	calibrationValue, _ := strconv.Atoi(values[0] + values[len(values)-1])
	return calibrationValue
}
