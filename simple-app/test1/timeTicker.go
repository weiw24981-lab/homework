package main

import (
	"fmt"
	"time"
)

func TimeDemo() {
	nowTime := time.Now()
	fmt.Printf("当前时间: %v\n", nowTime)
	// 格式化时间
	fmt.Printf("格式化时间: %v\n", nowTime.Format("2006-01-02 15:04:05 "))

	ti, err := time.Parse("2006-01-02", "2026-04-04")
	if err != nil {
		fmt.Printf("时间解析失败: %v\n", err)
	}
	fmt.Printf("解析后的时间: %v\n", ti)

	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Printf("加载时区失败: %v\n", err)
	}
	ti2, err := time.ParseInLocation("2006-01-02 15:04:05", "2026-04-04 12:08:16", loc)
	if err != nil {
		fmt.Printf("时间解析失败: %v\n", err)
	}
	fmt.Printf("解析后的时间: %v\n", ti2)
}

func TickerDemo() {
	ticker := time.NewTicker(2000 * time.Millisecond)
	defer ticker.Stop()
	for i := 0; i < 5; i++ {
		<-ticker.C
		fmt.Printf("次数%d:", i)
		fmt.Printf("tick at %v\n", time.Now())
	}

}
