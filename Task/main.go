// package main

// import (
// 	"context"
// 	"fmt"
// 	"os"
// 	"sync"
// 	"time"
// )

// func main() {

// 	fmt.Println("Multiple Goroutines")

// 	for i := 1; i <= 3; i++ {
// 		go func(num int) {
// 			fmt.Println("Goroutine:", num)
// 		}(i)
// 	}

// 	time.Sleep(1 * time.Second)

// 	fmt.Println("\n Producer-Consumer Pipeline")

// 	ch := make(chan int)

// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			fmt.Println("Produced:", i)
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	for value := range ch {
// 		fmt.Println("Consumed:", value)
// 	}

// 	fmt.Println("\n Concurrent File Processing")

// 	files := []string{"file1.txt", "file2.txt", "file3.txt"}

// 	var wg sync.WaitGroup

// 	for _, file := range files {
// 		wg.Add(1)

// 		go func(filename string) {
// 			defer wg.Done()

// 			os.WriteFile(filename, []byte("Hello Go"), 0644)

// 			data, _ := os.ReadFile(filename)

// 			fmt.Println("Read from", filename, ":", string(data))
// 		}(file)
// 	}

// 	wg.Wait()

// 	fmt.Println("\nShared Counter with Mutex")

// 	counter := 0
// 	var mutex sync.Mutex
// 	wg = sync.WaitGroup{}

// 	for i := 1; i <= 5; i++ {
// 		wg.Add(1)

// 		go func() {
// 			defer wg.Done()

// 			mutex.Lock()
// 			counter++
// 			mutex.Unlock()
// 		}()
// 	}

// 	wg.Wait()

// 	fmt.Println("Final Counter:", counter)

// 	fmt.Println("\nWorker Pool with Cancellation")

// 	jobs := make(chan int)

// 	ctx, cancel := context.WithCancel(context.Background())

// 	worker := func(id int) {
// 		for {
// 			select {

// 			case <-ctx.Done():
// 				fmt.Println("Worker", id, "stopped")
// 				return

// 			case job := <-jobs:
// 				fmt.Println("Worker", id, "processing job", job)
// 				time.Sleep(500 * time.Millisecond)
// 			}
// 		}
// 	}

// 	for i := 1; i <= 2; i++ {
// 		go worker(i)
// 	}

// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			jobs <- i
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)

// 	cancel()

// 	time.Sleep(1 * time.Second)
// }

package main

import (
	"fmt"
	"sync"
)

func printOdd(oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 10; i += 2 {
		<-oddChan

		fmt.Println("Odd:", i)

		evenChan <- true
	}
}

func printEven(oddChan chan bool, evenChan chan bool, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 2; i <= 10; i += 2 {
		<-evenChan

		fmt.Println("Even:", i)

		oddChan <- true
	}
}

func main() {
	var wg sync.WaitGroup

	oddChan := make(chan bool)
	evenChan := make(chan bool)

	wg.Add(2)

	go printOdd(oddChan, evenChan, &wg)
	go printEven(oddChan, evenChan, &wg)

	oddChan <- true

	wg.Wait()
}
