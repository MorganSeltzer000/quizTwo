package main

import (
	"fmt"
	"math"
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
	timeSum := make([]int64, 2, 2)
	for j := 0; j < 10; j++ {
		randIntsSlice := make([]int, totalNums, totalNums) //need to create 2 since sorting changes the elements
		randIntsStable := make([]int, totalNums, totalNums)
		for i := 0; i < totalNums; i++ {
			randIntsSlice[i] = rand.Intn(20)
		}
		copy(randIntsStable, randIntsSlice) //first arg is dst
		start := time.Now()
		sort.Slice(randIntsSlice, regSort)
		end := time.Since(start)
		timeSum[0] += end.Microseconds()
		fmt.Printf("It took %v to slice sort\n", end)
		start = time.Now()
		sort.SliceStable(randIntsStable, regSort)
		end = time.Since(start)
		timeSum[1] += end.Microseconds()
		fmt.Printf("It took %v to stable sort\n", end)
	}
	fmt.Printf("It took %v slice average, %v stable average, ratio %v\n",
		timeSum[0]/10, timeSum[1]/10, timeSum[0]/10/timeSum[1]/10)
	complexityOne := float64(timeSum[0]) / 10 / math.Log2(float64(totalNums))
	complexityTwo := float64(timeSum[1]) / 10 / math.Log2(float64(totalNums))
	fmt.Printf("After dividing by time complexity, %v slice and %v stable, ratio %v\n",
		complexityOne, complexityTwo, complexityOne/complexityTwo)
}
