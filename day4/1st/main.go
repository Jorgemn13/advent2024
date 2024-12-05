package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	textRows := []string{}

	for scanner.Scan() {
		row := scanner.Text()
		textRows = append(textRows, row)
	}

	answer := 0

	for rowIndex, textRow := range textRows {
		for i := 0; i < len(textRow); i++ {
			char := textRow[i]
			if char == 'X' {
				answer = answer + searchAround(rowIndex, i, textRows)
			}
		}
	}

	fmt.Println(answer)
}

func searchAround(rowIndex int, charIndex int, textRows []string) int {
	var directions = [8][2]int{
		{-1, 0},  // up
		{1, 0},   // down
		{0, -1},  // left
		{0, 1},   // right
		{-1, -1}, // up-left diagonal
		{-1, 1},  // up-right diagonal
		{1, -1},  // down-left diagonal
		{1, 1},   // down-right diagonal
	}
	counter := 0

	word := "XMAS"
	for _, dir := range directions {
		found := true
		for i := 0; i < len(word); i++ {
			newRow := rowIndex + i*dir[0]
			newCol := charIndex + i*dir[1]

			if newRow < 0 || newRow >= len(textRows) || newCol < 0 || newCol >= len(textRows[0]) {
				found = false
				break
			}

			if textRows[newRow][newCol] != word[i] {
				found = false
				break
			}
		}
		if found {
			counter++
		}
	}
	return counter
}
