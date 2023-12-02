package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseNumber(line string) int {
	re := regexp.MustCompile(`\d`)
	allDigits := re.FindAll([]byte(line), -1)

	a := int(allDigits[0][0]) - 48
	b := int(allDigits[len(allDigits)-1][0]) - 48

	return a*10 + b
}

var values = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9}

func getValue(barr string) int {
	if len(barr) == 1 {
		var d, err = strconv.Atoi(barr)
		if err == nil {
			return d
		} else {
			return 0
		}
	} else {
		return values[barr]
	}
}

func parseNumber2(line string) int {
	re := regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)

	first := re.FindString(line)
	var last string

	for i := len(line) - 1; i >= 0; i-- {
		slice := line[i:]
		last = re.FindString(slice)

		if last != "" {
			break
		}
	}

	if last == "" {
		last = first
	}

	a := getValue(first)
	b := getValue(last)

	return a*10 + b
}

func main() {

	// i := parseNumber(`sq5fivetwothree1`)
	// i := parseNumber2(`twone`)
	// fmt.Print(i)

	readFile, err := os.Open("input")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(readFile)
	scanner.Split(bufio.ScanLines)

	part1 := 0
	part2 := 0

	for scanner.Scan() {
		line := scanner.Text()
		part1 += parseNumber(line)
		part2 += parseNumber2(line)
	}

	fmt.Printf("part1: %d\n", part1)
	fmt.Printf("part2: %d\n", part2)
}
