import (
package main

import (
	"fmt"
sync "sync"
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



package main

	"fmt"
	"sync"
	"time"
)

var tasks = []func(){
	task1,
	task2,
	task3,
	task4,
}

const clockCycles = 10 // Number of clock cycles to perform tasks

func main() {
	for i := 0; i < clockCycles; i++ {
		start := time.Now()

		var wg sync.WaitGroup
		taskCh := make(chan func(), len(tasks))

		// Execute each task in parallel
		for _, task := range tasks {
			taskCh <- task
		}
		close(taskCh)

		// Process tasks concurrently
		for task := range taskCh {
			wg.Add(1)
			go executeTask(task, &wg)
		}

		// Wait for all tasks to complete
		wg.Wait()

		elapsed := time.Since(start)
		fmt.Printf("Tasks completed in %v\n", elapsed)
	}

	// Additional code outside main function
}

func executeTask(task func(), wg *sync.WaitGroup) {
	defer wg.Done()
	task()
}

func task1() {
	// Task 1 implementation
}

func task2() {
	// Task 2 implementation
}

func task3() {
	// Task 3 implementation
}

func task4() {
	// Task 4 implementation
}

