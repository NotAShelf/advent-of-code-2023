package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	Red   = "red"
	Green = "green"
	Blue  = "blue"
)

func main() {
	pFlags := flag.String("p", "", "Choose the part of the puzzle to run. Valid values are: '1' or '2'")
	flag.Parse()

	if *pFlags == "" {
		fmt.Println("Please provide a valid flag: -p 1 or -p 2")
		return
	}

	if *pFlags == "1" {
		part1()
	} else if *pFlags == "2" {
		part2()
	} else {
		fmt.Println("Invalid flag provided. Please use -p 1 or -p 2")
	}
}

func openFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	return bufio.NewScanner(file)
}

func part1() {
	file, err := os.Open("cubes.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	s := 0
	actualGame := map[string]int{
		Red:   12,
		Green: 13,
		Blue:  14,
	}

	for scanner.Scan() {
		line := scanner.Text()
		gameIndex, _ := strconv.Atoi(strings.Split(strings.Split(line, " ")[1], ":")[0])
		game := map[string]int{
			Red:   0,
			Green: 0,
			Blue:  0,
		}
		turns := strings.Replace(strings.TrimSpace(strings.Split(line, ":")[1]), ";", ",", -1)
		cubes := strings.Split(turns, ",")
		for _, cube := range cubes {
			cubeParts := strings.Split(strings.TrimSpace(cube), " ")
			number, _ := strconv.Atoi(cubeParts[0])
			color := cubeParts[1]
			if game[color] < number {
				game[color] = number
			}
		}
		if checkGame(game, actualGame) {
			s += gameIndex
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	fmt.Println(s)
}

func part2() {
	file, err := os.Open("cubes.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	s := 0

	for scanner.Scan() {
		line := scanner.Text()
		game := map[string]int{
			Red:   0,
			Green: 0,
			Blue:  0,
		}
		turns := strings.Replace(strings.TrimSpace(strings.Split(line, ":")[1]), ";", ",", -1)
		cubes := strings.Split(turns, ",")
		for _, cube := range cubes {
			cubeParts := strings.Split(strings.TrimSpace(cube), " ")
			number, _ := strconv.Atoi(cubeParts[0])
			color := cubeParts[1]
			if game[color] < number {
				game[color] = number
			}
		}
		s += game[Red] * game[Green] * game[Blue]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file.Close()
	fmt.Println(s)
}

func checkGame(game, actualGame map[string]int) bool {
	for color := range actualGame {
		if actualGame[color] < game[color] {
			return false
		}
	}
	return true
}
