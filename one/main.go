package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := os.Open("one/input.txt")
	if err != nil {
		fmt.Println("An error occured while opening the input file")
		panic(err);
	}

	scanner := bufio.NewScanner(file)

	// Digit regular expression definition - \d to match single digits and not the whole character group
	digitRe, err := regexp.Compile(`\d`); if err != nil {
		fmt.Println("Failed to compile regular expression for finding digits")
		panic(err)
	}

	calibrationTotal := 0
	// Using scanner.Scan reads each line of the opened file.
	// Note the entire file is not loaded into memory here.
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("Calibrating New Line: %s\n", line)
		matches := digitRe.FindAllString(line, -1)
		// Per AoC, the first and last digits are read as the first occurence of a digit forwards and backwards
		// So, if there's only one digit, that digit is considered both the first and last occurence.
		if len(matches) < 1 {
			fmt.Println("No Matches Found in this Line")
			continue
		}

		first := matches[0]
		last := matches[len(matches) - 1]
		calibrationStr := first + last


		fmt.Printf("Found a total of %d digits in the string. %s and %s are the first and last digits\n", len(matches), first, last)
		fmt.Printf("Calibration String is %s\n", calibrationStr)
		
		calibration, err := strconv.Atoi(calibrationStr); if err != nil {
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