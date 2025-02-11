package main

import (
	"errors"
	"fmt"
)

func main() {
	var err error
	if err == nil {
		err = errors.New("This error generating")
		fmt.Print("There was an error")
	} else {
		fmt.Print("There was no error")
	}
}
