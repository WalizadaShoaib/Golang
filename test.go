package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)
	go birds("parrot")
	go mammals("sheep")

	wg.Wait()
	fmt.Println("OS", runtime.GOOS)
	fmt.Println("Number of CPUs", runtime.NumCPU())
	fmt.Println("Arch", runtime.GOARCH)
	fmt.Println("Goroutine", runtime.NumGoroutine())

}

func printTime() {
	fmt.Println(time.Now().Format("15:04:05"))
}

func birds(name string) {
	for i := 1; i <= 10; i++ {

		fmt.Println(i, name)
		printTime()

	}
	wg.Done()
}

func mammals(name string) {
	for i := 1; i <= 10; i++ {
		fmt.Println(i, name)
		printTime()
	}
	wg.Done()
}
