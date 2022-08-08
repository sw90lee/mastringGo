package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {
	n := flag.Int("n", 20, "고루틴을 위한 숫자를 넣어주세요")
	flag.Parse()
	count := *n
	fmt.Printf("Going to create %d goroutines. \n", count)

	var waitGroup sync.WaitGroup

	fmt.Printf("%#v\n", waitGroup)
	for i := 0; i < count; i++ {
		waitGroup.Add(1)
		go func(x int) {
			defer waitGroup.Done()
			fmt.Printf("%d ", x)
		}(i)
	}

	fmt.Printf("%#v\n", waitGroup)
	waitGroup.Wait()
	fmt.Println("\nExiting...")

}
