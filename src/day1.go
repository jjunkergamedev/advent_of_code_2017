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

func main() {
	// f, err := os.Open("day1_input.txt")
	// check(err)

	dat, err := ioutil.ReadFile("day1_input.txt")
	check(err)
	fmt.Println(string(dat))

	fmt.Println("Calculating Checksum...")

	// prev := strconv.ParseInt(dat[len(dat)-1])
	// fmt.Println("prev: ", prev)
	sum := 0
	prev := dat[len(dat) - 1]
	for i, s := range dat {
		if(s == prev) {
			fmt.Println("i: ", i, " s: ", num(s), "prev: ", num(prev))
			sum += num(s)
		}
		prev = s
	}

	fmt.Println("sum: ", sum)
}
