package main

import (
	"fmt"
	"testing"
)

func TestExampleCleanup(t *testing.T) {
	fmt.Println("SETUP")
	t.Cleanup(func() {
		fmt.Println("TEARDOWN ON CLEANUP")
	})

	t.Run("First", func(t *testing.T) {
		fmt.Println("ok")
	})

	t.Run("Second", func(t *testing.T) {
		fmt.Println("ok")
	})

	t.Run("Third", func(t *testing.T) {
		// t.Fatal("falat test")
		panic("i am panuc!")
	})

	fmt.Println("TEARDOWN AT END")
}
