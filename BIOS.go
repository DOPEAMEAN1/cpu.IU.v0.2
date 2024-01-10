package main

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("Task 1")
}

func task2() {
	fmt.Println("Task 2")
}

func task3() {
	fmt.Println("Task 3")
}

func task4() {
	fmt.Println("Task 4")
}

func main() {
	tasks := []func(){
		task1,
		task2,
		task3,
		task4,
	}

	clockCycles := 10 // Number of clock cycles to perform tasks

	for i := 0; i < clockCycles; i++ {
		start := time.Now()

		// Execute each task in parallel
		for _, task := range tasks {
			go task()
		}

		// Wait for all tasks to complete
		time.Sleep(time.Millisecond * 250)

		elapsed := time.Since(start)
		fmt.Printf("Tasks completed in %v\n", elapsed)
	}
}
```