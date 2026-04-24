package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func GmpDemo() {
	start := time.Now()
	fmt.Printf("当前系统的CPU数量: %d\n", runtime.NumCPU())
	prev := runtime.GOMAXPROCS(0)
	fmt.Printf("当前的GOMAXPROCS值: %d\n", prev)

	// 设置GOMAXPROCS为2
	runtime.GOMAXPROCS(2)
	fmt.Printf("新的GOMAXPROCS值: %d\n", runtime.GOMAXPROCS(0))

	var wg sync.WaitGroup
	taskCount := 10
	wg.Add(taskCount)

	for i := 0; i < taskCount; i++ {
		go func(taskID int) {
			defer wg.Done()

			sum := 0
			for j := 0; j < 1000000; j++ {
				sum += j % (taskID + 1)

			}
			fmt.Printf("任务%d完成，结果: %d\n", taskID, sum)
		}(i)

	}
	wg.Wait()

	elapsed := time.Since(start)
	fmt.Printf("所有任务完成，耗时: %s\n", elapsed)
	fmt.Println("提示：运行时可配合命令 `GODEBUG=schedtrace=1000,scheddetail=1 go run gmp.go` 观察调度日志。")

	// 恢复原始的 GOMAXPROCS，避免影响其他程序
	runtime.GOMAXPROCS(prev)
}

// func main() {
// 	GmpDemo()
// }
