package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func partSum(nums []int, myChan chan int) {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	myChan <- sum
}

func main() {
	if len(os.Args) < 2 { //0 is name
		panic("No arg given, exiting")
	}
	totalNums, _ := strconv.Atoi(os.Args[1])
	randInts := make([]int, totalNums, totalNums)
	for i := 0; i < totalNums; i++ {
		randInts[i] = rand.Intn(20)
	}
	start := time.Now()
	sum := 0
	for i := 0; i < totalNums; i++ {
		sum += randInts[i]
	}
	end := time.Since(start)
	fmt.Printf("It took %v to sum the slice in main, and got %d\n", end, sum)
	start = time.Now()
	myChan := make(chan int)
	go partSum(randInts[:totalNums/2], myChan) //sum each half in parallel
	go partSum(randInts[totalNums/2:], myChan)
	sum = (<-myChan) + (<-myChan)
	end = time.Since(start)
	fmt.Printf("It took %v to sum the slice in parallel, and got %d\n", end, sum)
}
