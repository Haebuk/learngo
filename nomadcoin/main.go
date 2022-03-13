package main

import (
	"fmt"
	"time"
)

func countToTen(c chan int) {
	for i := range [10]int{} {
		c <- i
		time.Sleep(1 * time.Second)
	}
}

func main() {
	c := make(chan int)
	go countToTen(c)
	for {
		a := <-c
		fmt.Println(a)
	}
} 