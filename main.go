package main

import (
	"fmt"
)

func main() {
	task := new(Task)
	var f func(id string)
	f = func(i string) {
		fmt.Printf(i + "\n")
	}
	task.Init(1, 2, f)
	task.PutQueue("1")
	task.PutQueue("2")
	task.Start()
	for {

	}
}
