package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// ctx := context.Background()
	// context.TODO()
	// ctx, cancel := context.WithCancel(ctx)
	// cancel()

	ctx := context.Background()
	duration := 1500 * time.Millisecond
	// d := time.Now().Add(duration)
	// ctx, cancel := context.WithCancel(ctx)
	// ctx, cancel := context.WithDeadline(ctx, d)
	ctx = context.WithValue(ctx, "test", "hello")
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	/*go func() {
		err := cancelRequest(ctx)
		if err != nil {
			cancel()
		}
	}()*/
	doRequest(ctx, "https://ya.ru")
}

func cancelRequest(ctx context.Context) error {
	time.Sleep(1500 * time.Millisecond)
	return fmt.Errorf("fail request")
}

func doRequest(ctx context.Context, requestStr string) {
	req, _ := http.NewRequest(http.MethodGet, requestStr, nil)
	req = req.WithContext(ctx)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
	select {
	case <-time.After(500 * time.Millisecond):
		fmt.Printf("Response completed, status code = %d\n", res.StatusCode)
		fmt.Println(ctx.Value("test"))
	case <-ctx.Done():
		fmt.Println("Request too long")
	}
}
