package main

import "fmt"

func main() {
	var el [3]int32
	// var ele2 []int32 = [3]int32{5, 7, 8}
	var elements = [...]int32{1, 2, 3}
	fmt.Println(el[0])
	fmt.Println(elements)

	var more = []int32{4, 3, 5}
	fmt.Printf("length %v, capacity %v\n", len(more), cap(more))
	more = append(more, 3)
	fmt.Printf("length %v, capacity %v", len(more), cap(more))
	fmt.Println(more)

	var more2 = []int32{10, 11}
	more2 = append(more2, more...)
	fmt.Println(more2)

	// var more3 []int32 = make([]int32, 3, 4)

	// var object map[string]int32= make(map[string]int32)
	var names map[string]int32 = map[string]int32{"gabin": 4, "ishimwe": 10}
	fmt.Println(names["xx"])

	var val, ok = names["gabin"]
	if ok {
		fmt.Printf("Value found %v", val)
	} else {
		fmt.Println("Value not found")
	}

	// delete value in map
	// delete(names, "gabin")
	// fmt.Print(names)

	// iterate over map
	for name, age := range names {
		fmt.Printf("name %v, age %v", name, age)
	}

	// iterate over array
	for i, v := range more2 {
		fmt.Printf("name %v, age %v\n", i, v)
	}

	var i int = 0
	for i < 3 {
		fmt.Println("It is less than i")
		i++
	}

	for j := 0; j < 10; j++ {
		fmt.Println(j)
	}
}
