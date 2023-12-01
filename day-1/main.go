package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func isDigit(char string) bool {
	matched, _ := regexp.MatchString(`\d`, char)
	return matched
}

func formDigit(value string) int {
	var left, right string

	for _, char := range value {
		if isDigit(string(char)) {
			left = string(char)
			break
		}
	}

	for i := len(value) - 1; i >= 0; i-- {
		if isDigit(string(value[i])) {
			right = string(value[i])
			break
		}
	}

	leftDigit, _ := strconv.Atoi(left)
	rightDigit, _ := strconv.Atoi(right)
	return leftDigit*10 + rightDigit
}

func sum(a, b int) int {
	return a + b
}

func main() {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "1abc2", expected: 12},
		{input: "pqr3stu8vwx", expected: 38},
		{input: "a1b2c3d4e5f", expected: 15},
	}

	for _, testCase := range testCases {
		result := formDigit(testCase.input)
		fmt.Printf("sum(%s) => %d\n", testCase.input, result)
	}

	testcase := "1abc2\n" + "pqr3stu8vwx\n" + "a1b2c3d4e5f\n" + "treb7uchet"
	lines := strings.Split(testcase, "\n")
	total := 0

	for _, line := range lines {
		total += formDigit(line)
	}

	fmt.Println("\nresult should be 142:", total)

	file, err := os.Open("calibration.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	lines = strings.Split(string(content), "\n")
	total = 0

	for _, line := range lines {
		total += formDigit(line)
	}

	fmt.Println("Sum of all calibration values:", total)
}
