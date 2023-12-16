package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func GetDigitsFromString(re *regexp.Regexp, input string, output []string) ([]string, error) {
	if output == nil {
		output = make([]string, 0)
	}
	if len(input) == 0 {
		return output, nil
	}

	if len(input) == 1 {
		match := re.FindString(input)
		if match != "" {
			output = append(output, match)
		}
		return output, nil
	}
	matchIndex := re.FindStringIndex(input)
	if matchIndex == nil {
		return output, nil
	}
	match := input[matchIndex[0]:matchIndex[1]]
	if match != "" {
		output = append(output, match)
	}
	return GetDigitsFromString(re, input[matchIndex[0]+1:], output)
}

func main() {
	file, err := os.Open("one/input.txt")
	if err != nil {
		fmt.Println("An error occured while opening the input file")
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	// Create a Map of word digits and their int values to use in a lookup down the line
	wordDigits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	wordDigitMap := make(map[string]int)
	for i, v := range wordDigits {
		wordDigitMap[v] = i + 1
	}

	re, err := regexp.Compile(`\d|one|two|three|four|five|six|seven|eight|nine`)
	if err != nil {
		fmt.Println("Failed to compile regular expression for finding digits")
		panic(err)
	}
	// Digit regular expression definition - \d to match single digits and not the whole character group
	digitOnlyRe, err := regexp.Compile(`\d`)
	if err != nil {
		fmt.Println("Failed to compile digit only regex")
		panic(err)
	}

	calibrationTotal := 0
	// Using scanner.Scan reads each line of the opened file.
	// Note the entire file is not loaded into memory here.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Calibrating New Line: %s\n", line)
		matches, err := GetDigitsFromString(re, line, nil)
		if err != nil {
			fmt.Println("Failed to Get Digits from String")
			panic(err)
		}
		// Per AoC, the first and last digits are read as the first occurence of a digit forwards and backwards
		// So, if there's only one digit, that digit is considered both the first and last occurence.
		if len(matches) < 1 {
			fmt.Println("No Matches Found in this Line")
			continue
		}

		var first = matches[0]
		var last = matches[len(matches)-1]

		if digitOnlyRe.MatchString(first) == false {
			// String is not a digit so must be a word digit
			first = strconv.Itoa(wordDigitMap[first])
		}
		if digitOnlyRe.MatchString(last) == false {
			last = strconv.Itoa(wordDigitMap[last])
		}
		calibrationStr := first + last

		fmt.Printf("Found a total of %d digits in the string. %s and %s are the first and last digits\n", len(matches), first, last)
		fmt.Printf("Calibration String is %s\n", calibrationStr)

		calibration, err := strconv.Atoi(calibrationStr)
		if err != nil {
			fmt.Printf("Failed to convert calibration string: %s to an integer\n", calibrationStr)
			panic(err)
		}
		calibrationTotal += calibration
		fmt.Printf("Line calibration completed. Input: %s; Calibration: %d; TotalCalibration: %d\n\n", line, calibration, calibrationTotal)
	}

	if scanner.Err(); err != nil {
		fmt.Println("An error occured while reading content of the file")
		panic(err)
	}

	fmt.Printf("Calibration Complete. Calibration Value = %d\n\n", calibrationTotal)

	defer file.Close()
}
