package main

import (
	"fmt"
	// "os"
	"io/ioutil"
	"strings"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v 
}

func part1(data string) {
	fmt.Println("Part 1: Calculating Checksum...")

	var sum int = 0

	var lines []string = strings.Split(data, "\n")
	for _, s := range lines {
		fmt.Println(s)
		var nums []string = strings.Split(s, "\t")

		min, _ := strconv.Atoi(nums[0])
		var max int = min
		for i := 1; i < len(nums); i++ {
			n, _ := strconv.Atoi(nums[i])
			if n < min {
				min = n
			} else if n > max {
				max = n
			}
		}
		diff := Abs(max - min)
		fmt.Println("diff: ", diff)
		sum += diff
	}


	fmt.Println("Part 1 sum: ", sum)
}

func part2(data string) {
	fmt.Println("Part 2: Calculating Checksum...")

	var sum int = 0

	var row []int

	var lines []string = strings.Split(data, "\n")
	for rowNum, s := range lines {
		fmt.Println(s)
		var nums []string = strings.Split(s, "\t")

		// find the only 2 nums that divide evenly into each other.
		// sum the result.

		// brute force:
		// for each num
		// 		for each other num2
		//			if num evenly divides num2, use it and continue

		// alternate:
		// for each row
		//		sort the row
		//		for n from highest to lowest
		//			for m from n - 1 to lowest
		//				if n evenly divides m, use it and continue

		// another alternate:
		// build a list of "unique" numbers? what if a number appears twice, is 2/2 = 1 the answer for that row?
		// so, as we convert the numbers, add them to a list that we build
		// and, check all the previously converted numbers to see if they divide either way
		row = row[:0]
		rowResult := 0
		for i := 0; i < len(nums) && rowResult == 0; i++ {
			curNum, _ := strconv.Atoi(nums[i])

			for _, n := range row {
				if n % curNum == 0 {
					rowResult = n / curNum
					fmt.Println(n, " divided by ", curNum)
					break
				} else if curNum % n == 0 {
					rowResult = curNum / n
					fmt.Println(curNum, " divided by ", n)
					break
				}
			}

			row = append(row, curNum)
		}

		fmt.Println("Row ", rowNum, " result=", rowResult)
		sum += rowResult
	}

	fmt.Println("Part 2 sum: ", sum)
}

func main() {
	dat, err := ioutil.ReadFile("../data/day2_input_1.txt")
	check(err)
	fmt.Println(string(dat))

	part1(string(dat))

	dat, err = ioutil.ReadFile("../data/day2_input_2.txt")
	check(err)
	fmt.Println(string(dat))

	part2(string(dat))
}
