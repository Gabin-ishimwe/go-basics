package main

import (
	"fmt"
	"strings"
)

type teacher struct {
	firstName string
	lastName  string
}

type studentsProfile struct {
	name  string
	age   int8
	score int32
	teacher
}

type score interface {
	studentScore() int32
}

func main() {
	var myName = "gabin"
	var resume = "resume"
	fmt.Println(len(myName))
	fmt.Println(len(resume))

	for i, v := range resume {
		fmt.Printf("index %v, value %v", i, v)
	}

	var chars = []string{"g", "a", "b", "i", "n"}
	var results = ""
	var res strings.Builder
	for i := range chars {
		results += chars[i]
		res.WriteString(chars[i])
	}
	fmt.Printf("formatted name %v => ", results)
	fmt.Printf("new results %v ", res.String())

	// Student function
	studentProps()
}

// Interesting...
func (student studentsProfile) studentScore() int32 {
	return student.score * 2
}

func calculateScore(method score, percentage int32) int32 {
	return method.studentScore() * percentage
}

func studentProps() {
	var myStudent = studentsProfile{name: "ishimwe", age: 4, score: 10, teacher: teacher{firstName: "first", lastName: "last"}}
	fmt.Println(myStudent)
	fmt.Println(myStudent.age)
	fmt.Println(myStudent.teacher)

	var testing = struct {
		first string
		last  string
	}{"gabin", "ishimwe"}
	fmt.Println(testing)

	fmt.Println(myStudent.studentScore())

	fmt.Println(calculateScore(myStudent, 5))
}
