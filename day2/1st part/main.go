package main1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main1() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var numSlice = [][]int{}

	for scanner.Scan() {
		row := scanner.Text()

		numStrings := strings.Split(row, " ")

		var convertedNums = []int{}
		for i := 0; i < len(numStrings); i++ {
			convertedNum, err := strconv.Atoi(numStrings[i])
			if err != nil {
				log.Fatal(err)
			}
			convertedNums = append(convertedNums, convertedNum)
		}

		if err != nil {
			log.Fatal(err)
		}

		numSlice = append(numSlice, convertedNums)
	}

	var differenceList = [][]int{}

	for i := 0; i < len(numSlice); i++ {
		differenceList = append(differenceList, []int{})
		for j := 0; j < len(numSlice[i])-1; j++ {
			num := numSlice[i][j]
			nextNum := numSlice[i][j+1]

			differenceList[i] = append(differenceList[i], nextNum-num)
		}
	}

	answer := 0
	for i := 0; i < len(differenceList); i++ {
		if checkSafety(differenceList[i]) == true {
			answer = answer + 1
		}
	}

	fmt.Print(answer)
}

func checkSafety(daList []int) bool {
	binSpottedDown := false
	binSpottedUp := false
	for _, item := range daList {
		if item == 0 || item > 3 || item < -3 {
			return false
		}
		if item > 0 {
			binSpottedUp = true
		}
		if item < 0 {
			binSpottedDown = true
		}
	}

	if binSpottedUp && binSpottedDown {
		return false
	}

	return true
}
