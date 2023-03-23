// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// // This program should go to 11, but it seemingly only prints 1 to 10.
// func main() {
// 	ch := make(chan int)
// 	wg := new(sync.WaitGroup)
// 	go Print(ch, wg)

// 	for i := 1; i <= 11; i++ {
// 		wg.Add(1)

// 		ch <- i
// 		wg.Wait()
// 	}
// 	close(ch)
// }

// // Print prints all numbers sent on the channel.
// // The function returns when the channel is closed.
// func Print(ch <-chan int, wg *sync.WaitGroup) {
// 	for n := range ch { // reads from channel until it's closed
// 		time.Sleep(10 * time.Millisecond) // simulate processing time

// 		wg.Done()
// 		fmt.Println(n)
// 	}
// }
