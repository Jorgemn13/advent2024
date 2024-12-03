package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	numRegex, err := regexp.Compile("mul\\([0-9]{1,3},[0-9]{1,3}\\)")
	if err != nil {
		log.Fatal(err)
	}

	var text string

	for scanner.Scan() {
		row := scanner.Text()
		text += row
	}

	allMatches := numRegex.FindAllString(text, -1)

	answer := 0
	for _, match := range allMatches {

		res := strings.Split(match, ",")
		firstNum, _ := strconv.Atoi(res[0][4:])
		secondNum, _ := strconv.Atoi(res[1][:len(res[1])-1])

		answer += firstNum * secondNum
	}

	fmt.Println(answer)
}
