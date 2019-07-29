package main

import (
	"fmt"
	"github.com/xpfyg/task"
)

func main() {
	fmt.Print("hell")
	task = new(task.Task)
	task.Init(1, 1, func(id sttring) {
		fmt.Printf("asdasdasd")
	}())
	task.PutQueue("asda")
	task.Start()
	for {

	}
}
