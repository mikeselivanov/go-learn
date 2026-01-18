package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("Quit cought")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	/*
		go func() {
			for i := range 10 {
				fmt.Printf("%v = %v\n", i, <-c)
			}
			quit <- 0
		}()
		fibonacci(c, quit)
	*/
	go fibonacci(c, quit)
	for i := range 20 {
		fmt.Printf("%v = %v\n", i, <-c)
	}
	quit <- 0
}
