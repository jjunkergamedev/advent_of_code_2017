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
