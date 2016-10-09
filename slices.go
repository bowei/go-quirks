package main

import "fmt"

func main() {
	a := [10]int{}
	for i := range a {
		a[i] = i
	}

	doAppend(a[:5])
	fmt.Printf("a = %v\n", a)

	b := [5]int{}
	for i := range b {
		b[i] = i
	}
	doAppend(b[:5])
	fmt.Printf("b = %v\n", b)
}

// Append may or may not affect the underlying array, depending on the need
// for allocation.
func sideEffectAppend(slice []int) []int {
	fmt.Printf("doAppend %v\n", slice)
	slice = append(slice, 10)
	slice[0] = 10
	return slice
}
