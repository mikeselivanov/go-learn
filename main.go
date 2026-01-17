package main

import "fmt"

const BigInt = 1 << 100
const SmallInt = BigInt >> 99

type Vertex struct {
	X int
	Y int
}

func main() {
	primes := []int{1, 3, 6, 7, 9, 23, 54}
	printSlice(primes)

	primes = primes[:0]
	printSlice(primes)

	// extend the length
	primes = primes[:4]
	printSlice(primes)

	primes = primes[2:]
	printSlice(primes)

	primes = primes[:5]
	printSlice(primes)
}

func printSlice(slice []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(slice), cap(slice), slice)
}
