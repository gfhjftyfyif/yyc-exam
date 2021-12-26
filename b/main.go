package main

import (
	"fmt"
)
func main() {
	var channel = make(chan string)
	a := "下山的路又堵起了"
	var b string
	go func() {
		channel <- a
	}()
	b = <- channel
	fmt.Println(b)
}