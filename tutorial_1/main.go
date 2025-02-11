package main

import (
	"errors"
	"fmt"
)

func main() {
	// declare variable
	var num int = 10
	fmt.Println(num)

	var age float32 = 10.5
	fmt.Println(age)

	var result = int(age) + (num)
	fmt.Println(result)

	var int1 int = 2
	var int2 int = 3

	fmt.Println(int1 / int2)
	fmt.Println(int1 % int2) // Get remainder

	var name string = "My name"
	var text string = `
Paragraph
	`
	fmt.Println(name)
	fmt.Println(text)

	var isOld bool = true
	fmt.Println(isOld)

	var var1, var2 = 10, 12

	var3, var4 := 11, 13

	fmt.Println(var1, var2)

	fmt.Println(var3, var4)

	const change bool = true
	fmt.Println(change)

	var first = "gabin"

	sample(first)

	// division
	var res, res2, err = division(11, 2)

	fmt.Println(res, res2)
	fmt.Printf("remainder %v", res2)

	if err == nil {
		fmt.Printf(err.Error())
	} else if err != nil {
		fmt.Println("It is not nil")
	} else {
		fmt.Println("It is none of the conditions")
	}
}

func sample(test string) {
	fmt.Println(test)
}

func division(first int, second int) (int, int, error) {
	var err error
	if first == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, nil
	}
	return (first / second), (first % second), err
}
