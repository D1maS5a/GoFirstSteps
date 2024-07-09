package main

import (
	"fmt"
	"sync"
	"time"
)

func myFunction() {
	fmt.Println("Выполнение в горутине")
	time.Sleep(time.Second)
	fmt.Println("Горутина выполнена")
}

func main() {
	fmt.Println("Старт основной горутины")
	go myFunction()
	fmt.Println("Основная горутина продолжает работу")

	time.Sleep(2 * time.Second)

	fmt.Println("Выполнение завершено")

	ch := make(chan int)
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
		close(ch)
	}()
	for num := range ch {
		fmt.Println(num)
	}

	var wg sync.WaitGroup

	for i := 1; i < 3; i++ {
		wg.Add(1) // Увеличение счетчика горутин в WaitGroup
		go worker(i, &wg)
	}
	wg.Wait()
	fmt.Println("ALL Gorutine complete")

	// Select
	ch2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 42
	}()
	select {
	case <-ch2:
		fmt.Println("Получено значение из канала")
	}
}

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Указываем, что горутина завершила свою работу
	fmt.Printf("Горутина %d завершила свою работу\n", id)
}
