package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	grid := [][]string{}

	for scanner.Scan() {
		row := scanner.Text()

		row2 := strings.Split(row, "")
		grid = append(grid, row2)
	}

	antennaPositions := findCommonAntennaPositions(grid)
	grid = placeAntinodes(grid, antennaPositions)

	answer := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] == "#" {
				answer++
			}
		}
	}

	printGrid(grid)
	fmt.Println(answer)
}

func placeAntinodes(grid [][]string, antennaPositions map[string][][]int) [][]string {
	antinodeSpots := [][]int{}
	for _, cords := range antennaPositions {
		for i := 0; i < len(cords); i++ {
			for j := i + 1; j < len(cords); j++ {
				diff := []int{cords[j][0] - cords[i][0], cords[j][1] - cords[i][1]}

				newY := cords[i][0]
				newX := cords[i][1]
				for newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) {
					antinodeSpots = append(antinodeSpots, []int{newY, newX})
					newY = newY - diff[0]
					newX = newX - diff[1]
				}

				newY = cords[i][0]
				newX = cords[i][1]
				for newY >= 0 && newY < len(grid) && newX >= 0 && newX < len(grid[0]) {
					antinodeSpots = append(antinodeSpots, []int{newY, newX})
					newY = newY + diff[0]
					newX = newX + diff[1]
				}
			}

		}
	}

	for _, spot := range antinodeSpots {
		if spot[0] >= 0 && spot[0] < len(grid) && spot[1] >= 0 && spot[1] < len(grid[0]) {
			grid[spot[0]][spot[1]] = "#"
		}
	}

	return grid
}

func findCommonAntennaPositions(grid [][]string) map[string][][]int {
	m := make(map[string][][]int)

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] != "." {
				m[grid[i][j]] = append(m[grid[i][j]], []int{i, j})
			}
		}
	}

	return m
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
	fmt.Println("------------------------")
}
