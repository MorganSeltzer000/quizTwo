package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func regSort(i, j int) bool { //didnt want to type twice
	return i < j
}

func main() {
	if len(os.Args) < 2 { //0 is name
		panic("No arg given, exiting")
	}
	totalNums, _ := strconv.Atoi(os.Args[1])
	randIntsSlice := make([]int, totalNums, totalNums) //need to create 2 since sorting changes the elements
	randIntsStable := make([]int, totalNums, totalNums)
	for i := 0; i < totalNums; i++ {
		randIntsSlice[i] = rand.Intn(20)
	}
	copy(randIntsStable, randIntsSlice) //first arg is dst
	start := time.Now()
	sort.Slice(randIntsSlice, regSort)
	end := time.Since(start)
	fmt.Printf("It took %v to slice sort\n", end)
	start = time.Now()
	sort.SliceStable(randIntsStable, regSort)
	end = time.Since(start)
	fmt.Printf("It took %v to stable sort\n", end)
}
