package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu sync.Mutex
	c  map[string]int
}

type CounterRW struct {
	mu sync.RWMutex
	c  map[string]int
}

func (c *Counter) Inc(key string) {
	c.mu.Lock()
	c.c[key]++
	c.mu.Unlock()
}

func (c *Counter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock() // нужен когда допустим логика if else с несколькими return-ами
	val := c.c[key]
	if val == 30 {
		return c.c[key] + 30
	}
	return c.c[key]
}

func main() {
	key := "test"
	c := Counter{c: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc(key)
	}

	time.Sleep(3 * time.Second)
	fmt.Println(c.Value(key))

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1) // 1 - сколько горутин запускаем (по сути является счетчиком) прибавляет 1
		go func() {
			defer wg.Done() // отнимает 1

			fmt.Printf("%d goroutine sleep... \n", i)
			time.Sleep(300 * time.Millisecond)
		}()
	}
	wg.Wait()
	fmt.Println("All done")

	var wg1 sync.WaitGroup
	var counter uint64
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		wg1.Add(1) // 1 - сколько горутин запускаем (по сути является счетчиком) прибавляет 1
		go func() {
			defer wg1.Done() // отнимает 1
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
				// atomic.AddUint64(&counter, 1) // потокобезопасное применение счетчика
			}
		}()
	}
	wg1.Wait()
	fmt.Printf("All done, counter %d \n", counter)

	var once sync.Once
	done := make(chan bool)

	for i := 0; i < 30; i++ {
		go func() {
			once.Do(func() {
				fmt.Println("the only one")
			})
			fmt.Println("the many")
			done <- true
		}()
	}
	for i := 0; i < 30; i++ {
		<-done
	}
}

func (c *CounterRW) CountMe() map[string]int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.c
}

func (c *CounterRW) CountMeAgain() map[string]int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.c
}
