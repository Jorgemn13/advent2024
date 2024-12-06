package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	pageOrderMap := make(map[string][]string)
	updates := [][]string{}

	for scanner.Scan() {
		row := scanner.Text()
		if strings.Contains(row, "|") {
			pages := strings.Split(row, "|")
			pageOrderMap[pages[0]] = append(pageOrderMap[pages[0]], pages[1])
		} else {
			nums := strings.Split(row, ",")
			updates = append(updates, nums)
		}
	}

	updates = updates[1:]
	answer := 0

	for i := 0; i < len(updates); i++ {
		update := updates[i]
		updateCopy := make([]string, len(update))
		copy(updateCopy, update)

		slices.SortFunc(update, func(i, j string) int {
			nums, ok := pageOrderMap[i]
			if ok && slices.Contains(nums, j) {
				return -1
			}
			nums, ok = pageOrderMap[j]
			if ok && slices.Contains(nums, i) {
				return 1
			}

			return 0
		})

		if reflect.DeepEqual(update, updateCopy) {
			num, _ := strconv.Atoi(update[len(update)/2])
			answer += num
		}
	}
	fmt.Println(answer)
}
