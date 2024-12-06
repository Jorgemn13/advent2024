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

	answer := 0

	for i := 0; i < (len(grid)); i++ {
		for j := 0; j < (len(grid[i])); j++ {

			if startingPoint[0] == i && startingPoint[1] == j {
				continue
			}

			finished := false
			direction := "^"
			counterQ := 0

			copyGrid := deepCopy(grid)

			copyStartingPoint := make([]int, len(startingPoint))
			copy(copyStartingPoint, startingPoint)

			copyGrid[i][j] = "X"

			for !finished {
				copyGrid, copyStartingPoint, finished = move1(copyStartingPoint, copyGrid, &answer, &direction, &counterQ)
			}
		}
	}

	fmt.Println(answer)
}

func deepCopy(src [][]string) [][]string {
	dst := make([][]string, len(src))
	for i := range src {
		dst[i] = make([]string, len(src[i]))
		copy(dst[i], src[i])
	}
	return dst
}

func printGrid(grid [][]string) {
	for i := 0; i < len(grid); i++ {
		fmt.Println(grid[i])
	}
	fmt.Println("-------------------")
}

func move1(startingPoint []int, grid [][]string, answer *int, direction *string, counter *int) ([][]string, []int, bool) {
	corX := startingPoint[1]
	corY := startingPoint[0]

	if *direction == "^" {
		corY = corY - 1
	}
	if *direction == ">" {
		corX = corX + 1
	}
	if *direction == "<" {
		corX = corX - 1
	}
	if *direction == "v" {
		corY = corY + 1
	}

	if corX < 0 || corX >= len(grid[0]) || corY < 0 || corY >= len(grid) {
		return grid, startingPoint, true
	}

	if grid[corY][corX] == "q" {
		*counter++
	}

	if grid[corY][corX] == "q" && (*counter > (len(grid) * len(grid[0]))) {
		*answer++
		return grid, startingPoint, true
	}

	if grid[corY][corX] == "#" || grid[corY][corX] == "X" {
		if *direction == "^" {
			corY = corY + 1
			*direction = ">"
			grid[corY][corX] = "q"
		} else if *direction == ">" {
			corX = corX - 1
			*direction = "v"
			grid[corY][corX] = "q"
		} else if *direction == "<" {
			corX = corX + 1
			*direction = "^"
			grid[corY][corX] = "q"
		} else if *direction == "v" {
			corY = corY - 1
			*direction = "<"
			grid[corY][corX] = "q"
		}
	}
	if grid[corY][corX] != "q" {
		grid[corY][corX] = *direction
	}

	startingPoint[0] = corY
	startingPoint[1] = corX

	return grid, startingPoint, false
}
