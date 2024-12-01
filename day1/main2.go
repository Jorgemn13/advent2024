package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var row string

	var leftArray = []int{}
	var rightArray = []int{}

	for scanner.Scan() {
		row = scanner.Text()
		bothNumbers := strings.Fields(row)

		leftNum, _ := strconv.Atoi(bothNumbers[0])
		rightNum, _ := strconv.Atoi(bothNumbers[1])

		leftArray = append(leftArray, leftNum)
		rightArray = append(rightArray, rightNum)
	}

	var count = 0
	var answer = 0

	for i := 0; i < len(leftArray); i++ {
		num := leftArray[i]
		for j := 0; j < len(rightArray); j++ {
			if num == rightArray[j] {
				count = count + 1
			}
		}

		answer = answer + (count * num)
		count = 0
	}

	fmt.Print(answer)

}
