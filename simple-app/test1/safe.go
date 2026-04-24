package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (c *SafeCounter) IncreamCnt(m int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
	fmt.Printf("协程%d调用increamCnt,目前计数为%d\n", m, c.count)
}

func GetCount(c *SafeCounter) int {
	return c.count
}

func NewSafeCounter() *SafeCounter {
	return &SafeCounter{}
}

func DoMutexDemo() {
	counter := NewSafeCounter()
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.IncreamCnt(i)
		}()
	}
	wg.Wait()
	fmt.Printf("最后计数为%d\n", GetCount(counter))
}

// 超时控制
func ContextTimeoutDemo() {
	cxt, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)
	defer cancel()
	ch := make(chan string)
	go func() {
		time.Sleep(3000 * time.Millisecond)
		ch <- "hello"
	}()

	select {
	case re := <-ch:
		fmt.Printf("收到结果:%s", re)
	case <-cxt.Done():
		fmt.Printf("超时:%s", cxt.Err())
	}

}

// 取消控制
func ContextCancelDemo() {
	cxt, cancel := context.WithCancel(context.Background())
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case <-cxt.Done():
				fmt.Printf("中止操作")
				return
			default:
				fmt.Printf("工作继续中:%d\n", i)
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()
	time.Sleep(1000 * time.Millisecond)
	cancel()

}

type SafeMap struct {
	mu   sync.RWMutex
	date map[string]int
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		date: make(map[string]int),
	}
}

func (s *SafeMap) ReadVebose(label string, key string, hold time.Duration) int {
	fmt.Printf("%s准备读取%s\n", label, key)
	s.mu.RLock()
	fmt.Printf("%s获得锁，准备读取%s\n", label, key)
	defer s.mu.RUnlock()
	if hold > 0 {
		time.Sleep(hold)
	}
	value, ok := s.date[key]
	if ok {
		fmt.Printf("%s读取结果:key%s=%d\n", label, key, value)
	} else {
		fmt.Printf("%s元素%s还未写入", label, key)
		value = -1
	}
	fmt.Printf("%s释放锁\n", label)
	return value
}

func (s *SafeMap) WriteVerbose(label, key string, value int, hold time.Duration) {
	fmt.Printf("%s准备写入 %s=%d\n", label, key, value)
	s.mu.Lock()
	fmt.Printf("%s获得锁，准备写入%s=%d\n", label, key, value)
	defer s.mu.Unlock()
	if hold > 0 {
		time.Sleep(hold)
	}
	s.date[key] = value
	fmt.Printf("%s写入完成，释放锁\n", label)
}

func (s *SafeMap) GetAll() map[string]int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	re := make(map[string]int)
	for k, v := range s.date {
		re[k] = v
	}
	return re
}

func OperateSafeMapDemo() {
	safeMap := NewSafeMap()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		safeMap.WriteVerbose("Write#1", "shared", 1, 1000*time.Millisecond)
	}()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			safeMap.ReadVebose(fmt.Sprintf("Read#%d", i), "shared", 1500*time.Millisecond)

		}(i)
	}
	wg.Wait()
	fmt.Printf("最终数据:%v\n", safeMap.GetAll())
}
