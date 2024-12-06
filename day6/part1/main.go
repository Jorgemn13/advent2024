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

	startingPoint := []int{0, 0}

	for i := 0; i < (len(grid)); i++ {
		for j := 0; j < (len(grid[i])); j++ {
			if grid[i][j] == "^" {
				startingPoint[0] = i
				startingPoint[1] = j
			}
		}
	}

	finished := false
	answer := 1

	for !finished {
		grid, startingPoint, finished = move1(startingPoint, grid, &answer)
	}

	fmt.Println(answer)

}

func move1(startingPoint []int, grid [][]string, answer *int) ([][]string, []int, bool) {
	corX := startingPoint[1]
	corY := startingPoint[0]
	direction := grid[corY][corX]

	if direction == "^" {
		corY = corY - 1
	}
	if direction == ">" {
		corX = corX + 1
	}
	if direction == "<" {
		corX = corX - 1
	}
	if direction == "v" {
		corY = corY + 1
	}

	if corX < 0 || corX >= len(grid[0]) || corY < 0 || corY >= len(grid) {
		return grid, startingPoint, true
	}

	if grid[corY][corX] == "#" {
		if direction == "^" {
			corY = corY + 1
			direction = ">"
		} else if direction == ">" {
			corX = corX - 1
			direction = "v"
		} else if direction == "<" {
			corX = corX + 1
			direction = "^"
		} else if direction == "v" {
			corY = corY - 1
			direction = "<"
		}
	}

	if grid[corY][corX] == "." {
		*answer++
	}

	grid[corY][corX] = direction
	startingPoint[0] = corY
	startingPoint[1] = corX

	return grid, startingPoint, false
}
