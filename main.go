package main

import "fmt"

func main() {
	ch := make(chan int)
	var count int

	go func() {
		ch <- 1
	}()

	go func() {
		count++
	}()

	<-ch

	fmt.Println(count)
}
