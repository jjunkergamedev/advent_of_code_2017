package main

import (
	"fmt"
	// "os"
	"io/ioutil"
	// "strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func num(v byte) int {
	return int(v - byte('0'));
}

func part1(data []byte) {
	fmt.Println("Part 1: Calculating Checksum...")

	sum := 0
	prev := data[len(data) - 1]
	for _, s := range data {
		if(s == prev) {
			// fmt.Println("i: ", i, " s: ", num(s), "prev: ", num(prev))
			sum += num(s)
		}
		prev = s
	}

	fmt.Println("Part 1 sum: ", sum)
}

func part2(data []byte) {
	fmt.Println("Part 2: Calculating Checksum...")

	sum := 0
	offset := len(data) / 2
	fmt.Println("offset: ", offset)
	for i, s := range data {
		other := data[(i + offset) % len(data)]
		if(s == other) {
			// fmt.Println("i: ", i, " s: ", num(s), "other: ", num(other))
			sum += num(s)
		}
	}

	fmt.Println("Part 2 sum: ", sum)
}

func main() {
	dat, err := ioutil.ReadFile("../data/day1_input.txt")
	check(err)
	fmt.Println(string(dat))

	part1(dat)

	dat, err = ioutil.ReadFile("../data/day1_input_2.txt")
	check(err)
	fmt.Println(string(dat))

	part2(dat)
}
