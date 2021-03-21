package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"test/queue"
	"time"
)

const (
	randMax = 100
)

func worker(gID, arrSize int) func() {
	//make random array
	arr := make([]int, arrSize)
	for i := 0; i < arrSize; i++ {
		arr[i] = rand.Intn(randMax)
	}
	insertionTime := time.Now()
	//return task
	return func() {
		//sort array
		sort.Ints(arr)
		//get output params
		min := arr[0]
		max := arr[len(arr)-1]
		mid := arr[arrSize/2]
		//print it
		fmt.Printf("%d %d %d %d %d\n", gID, insertionTime.UnixNano(), min, mid, max)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	//get cmd params
	writersPtr := flag.Int("writers", 2, "writers num")
	arrSizePtr := flag.Int("arr-size", 10, "array size")
	iterCountPtr := flag.Int("iter-count", 2, "iterations count")
	flag.Parse()
	//create new queue struct
	qu := queue.NewQueue()
	defer qu.Stop()
	wg.Add(*writersPtr)
	//start goroutines, put tasks
	for i := 0; i < *writersPtr; i++ {
		go func(gID int) {
			for k := 0; k < *iterCountPtr; k++ {
				qu.Put(worker(gID, *arrSizePtr))
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	//get and run tasks
	for {
		//get task
		task := qu.Get()
		if task == nil {
			break
		}
		//run task
		task()
	}
}
