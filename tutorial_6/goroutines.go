package main

import "fmt"

type lesson struct {
	title       string
	description string
}

type class struct {
	className        string
	classDescription string
}

type student[T lesson | class] struct {
	name   string
	sample T
}

func main() {
	var list1 = []int32{1, 2, 3}
	fmt.Println(genericSum(list1))

	var class1 = student[class]{
		name: "gabin",
		sample: class{
			className:        "pysics",
			classDescription: "gravity",
		},
	}

	var lesson1 = student[lesson]{
		name: "ishimwe",
		sample: lesson{
			title:       "math",
			description: "class",
		},
	}

	fmt.Printf("class1 %v", class1)
	fmt.Printf("lession1 %v", lesson1)

	var list2 = []float32{4, 5, 6}
	fmt.Println(genericSum(list2))
}

func genericSum[T float32 | int32](list []T) T {
	var sum T
	for _, v := range list {
		sum += v
	}
	return sum
}
