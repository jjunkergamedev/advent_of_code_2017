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

// func part2(data []byte) {
// 	fmt.Println("Part 2: Calculating Checksum...")

// 	sum := 0
// 	offset := len(data) / 2
// 	fmt.Println("offset: ", offset)
// 	for i, s := range data {
// 		other := data[(i + offset) % len(data)]
// 		if(s == other) {
// 			// fmt.Println("i: ", i, " s: ", num(s), "other: ", num(other))
// 			sum += num(s)
// 		}
// 	}

// 	fmt.Println("Part 2 sum: ", sum)
// }

func main() {
	dat, err := ioutil.ReadFile("../data/day2_input_1.txt")
	check(err)
	fmt.Println(string(dat))

	part1(string(dat))

	// dat, err = ioutil.ReadFile("day1_input_2.txt")
	// check(err)
	// fmt.Println(string(dat))

	// part2(dat)
}
