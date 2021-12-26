package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ch = make(chan string, 10)

func a(int) {
	var x string = "1"
	ch <- x
	wg.Done()
}
func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go a(i)
	}
	wg.Wait()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
