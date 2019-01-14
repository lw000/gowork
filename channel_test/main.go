// channel_test project main.go
package main

import (
	"fmt"
	"sync"
	"time"
)

func test() {

}

var (
	sharedIndex int = 0
	m           sync.Mutex
)

func worker(quit chan bool, tag int) {
	defer func() {
		fmt.Printf("leave[%d]\n", tag)
	}()

	// LEAVE:
	// 	for {
	// 		select {
	// 		case <-quit:
	// 			break LEAVE
	// 		default:
	// 			sharedIndex++
	// 			fmt.Printf("worker[%d] %d\n", tag, sharedIndex)
	// 		}
	// 	}

	for {
		select {
		case <-quit:
			return
		default:
			m.Lock()
			sharedIndex++
			j := sharedIndex
			m.Unlock()
			fmt.Printf("worker[%d] %d\n", tag, j)
		}
	}
}

func main() {
	// var message chan string = make(chan string)
	// go func(msg string) {
	// 	var data string
	// 	select {
	// 	case data = <-message:
	// 		{
	// 			fmt.Println("recv: " + data)
	// 		}
	// 		// default:
	// 		// 	fmt.Println("no communication")
	// 	}
	// }("pong")

	// go func(msg string) {
	// 	message <- msg
	// 	fmt.Println("start: " + msg)
	// }("ping")

	quit := make(chan bool)

	for i := 0; i < 10; i++ {
		go worker(quit, i)
	}

	time.AfterFunc(time.Duration(1)*time.Second, func() {
		// quit <- true
		close(quit)
	})

	time.Sleep(time.Duration(5) * time.Second)
}
