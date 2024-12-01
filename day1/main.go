package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main2() {
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

	leftMin := leftArray[0]
	rightMin := rightArray[0]

	answer := 0

	var leftIndex int
	var rightIndex int

	for len(leftArray) > 0 {
		for i := 0; i < len(leftArray); i++ {
			if leftArray[i] < leftMin {
				leftMin = leftArray[i]
				leftIndex = i
			}

			if rightArray[i] < rightMin {
				rightMin = rightArray[i]
				rightIndex = i
			}
		}

		answer = answer + int(math.Abs(float64(leftMin-rightMin)))
		if leftIndex == len(leftArray)-1 {
			leftArray = leftArray[:leftIndex]
		} else {
			leftArray = append(leftArray[:leftIndex], leftArray[leftIndex+1:]...)
		}

		if rightIndex == len(rightArray)-1 {
			rightArray = rightArray[:rightIndex]
		} else {
			rightArray = append(rightArray[:rightIndex], rightArray[rightIndex+1:]...)
		}

		if len(leftArray) > 0 {
			leftMin = leftArray[0]
			leftIndex = 0
		}
		if len(rightArray) > 0 {
			rightMin = rightArray[0]
			rightIndex = 0
		}
	}

	fmt.Print(answer)
}
