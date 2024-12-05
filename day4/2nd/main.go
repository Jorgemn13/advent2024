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
		if rowIndex == 0 || rowIndex == len(textRows)-1 {
			continue
		}

		for i := 1; i < len(textRow)-1; i++ {
			char := textRow[i]
			if char == 'A' {
				if searchAround(rowIndex, i, textRows) {
					answer++
				}
			}
		}
	}

	fmt.Println(answer)
}

func searchAround(rowIndex int, charIndex int, textRows []string) bool {

	return ((textRows[rowIndex-1][charIndex-1] == 'M' && textRows[rowIndex+1][charIndex+1] == 'S' || textRows[rowIndex-1][charIndex-1] == 'S' && textRows[rowIndex+1][charIndex+1] == 'M') && (textRows[rowIndex+1][charIndex-1] == 'M' && textRows[rowIndex-1][charIndex+1] == 'S' || textRows[rowIndex+1][charIndex-1] == 'S' && textRows[rowIndex-1][charIndex+1] == 'M'))

}
