# task
## 基于golang chan 实现的原生定时任务服务

### 用法
```
package main

import (
	"fmt"
	"github.com/xpfyg/task"
)

func main() {
	task := new(task.Task)
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
```