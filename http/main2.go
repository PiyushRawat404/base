// package main

// import "fmt"

// func greet(ch chan string) {
// 	ch <- "Hello from goroutine!" // Send value to channel
// }

// func main() {

// 	messageChannel := make(chan string)

// 	go greet(messageChannel)

// 	msg := <-messageChannel

// 	fmt.Println(msg)
// }