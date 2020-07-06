package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(3)

	go birds()
	go fishes()
	go human()
	wg.Wait()
	fmt.Println("Number of CPUs", runtime.NumCPU())

}

func timee() string {
	return time.Now().Format("15:04:05")

}

func birds() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Bird", i, timee())

	}
	wg.Done()
}

func fishes() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Fish", i, timee())

	}
	wg.Done()
}
func human() {
	for i := 1; i <= 100; i++ {
		fmt.Println("Human", i, timee())

	}
	wg.Done()
}
