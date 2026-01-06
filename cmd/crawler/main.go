package main

import (
	"fmt"
)

func main() {
	var consoleFlag ConsoleFlag
	err := ParseFlag(&consoleFlag)
	if err != nil {
		fmt.Println(err)
	}
}
