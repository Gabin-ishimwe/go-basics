package main

import "fmt"

func main() {
	fmt.Println("Learning pointers")
	var pointer *int32 = new(int32)
	// var pointers int32
	fmt.Printf("dereference pointer %v", *pointer)
	fmt.Printf("pointer %v", pointer)
	*pointer = 100

	var arr = [3]int8{1, 2, 5}

	fmt.Println(square(&arr))

	fmt.Println(arr)
}
func square(arr *[3]int8) [3]int8 {
	for i := range arr {
		arr[i] = arr[i] * arr[i]
	}
	return *arr
}
