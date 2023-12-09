package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		var first, second string

		for i := 0; i < len(runes); i++ {
			if unicode.IsDigit(runes[i]) {
				first = string(runes[i])
				break
			}
		}

		for i := len(runes) - 1; i >= 0; i-- {

			if unicode.IsDigit(runes[i]) {
				second = string(runes[i])
				break
			}
		}

		calibrationValue, _ := strconv.Atoi(first + second)
		sum += calibrationValue

	}

	fmt.Println("The sum is:", sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
