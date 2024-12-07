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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	numbers := make(map[int][]int)

	for scanner.Scan() {
		row := scanner.Text()

		nums := strings.Split(row, ":")
		keyString := nums[0]
		keyNum, _ := strconv.Atoi(keyString)
		values := strings.Split(strings.Join(nums[1:], ""), " ")

		var numValue []int
		for _, value := range values[1:] {
			val, err := strconv.Atoi(value)
			if err != nil {
				log.Fatal(err)
			}
			numValue = append(numValue, val)
		}
		numbers[keyNum] = numValue
	}

	answer := 0
	for goal, line := range numbers {
		if doesWork(goal, line) {
			answer += goal
		}
	}

	fmt.Println(answer)
}

func doesWork(goal int, values []int) bool {
	limit := powInt(3, len(values)-1)

	for i := 0; i < limit; i++ {
		result := values[0]
		hold := i
		for j := 1; j < len(values); j++ {
			if hold%3 == 0 {
				result += values[j]
			} else if hold%3 == 1 {
				result *= values[j]
			} else if hold%3 == 2 {
				result, _ = strconv.Atoi(fmt.Sprintf("%d%d", result, values[j]))
			}
			hold /= 3
		}
		if result == goal {
			return true
		}
	}

	return false
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}
