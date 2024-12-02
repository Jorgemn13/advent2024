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

	answer := 0
	for i := 0; i < len(numSlice); i++ {
		if checkSafety(numSlice[i]) {
			answer = answer + 1
		} else {
			for j := 0; j < len(numSlice[i]); j++ {
				modifiedList := make([]int, len(numSlice[i]))
				copy(modifiedList, numSlice[i])
				modifiedList = removeElement(modifiedList, j)
				if checkSafety(modifiedList) {
					answer = answer + 1
					break
				}
			}
		}
	}

	fmt.Println(answer)
}

func checkSafety(daList []int) bool {
	goodList := genSmallDifferenceList(daList)
	binSpottedDown := false
	binSpottedUp := false
	for _, item := range goodList {
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

func removeElement(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func genSmallDifferenceList(slice []int) []int {
	var differenceList []int

	for i := 0; i < len(slice)-1; i++ {
		num := slice[i]
		nextNum := slice[i+1]

		differenceList = append(differenceList, nextNum-num)
	}
	return differenceList
}
