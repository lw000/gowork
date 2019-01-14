// timer_test project main.go
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	timeout := time.After(time.Second * time.Duration(2))
	timeout1 := time.NewTimer(time.Second * time.Duration(3))
	go func() {
		for {
			select {
			case <-timeout:
				{
					fmt.Println("2s定时输出")
				}
			default:
				// time.Sleep(time.Millisecond * time.Duration(1))
			}

			select {
			case <-timeout1.C:
				{
					fmt.Println("3s定时输出")
				}
			default:
				// time.Sleep(time.Millisecond * time.Duration(1))
			}
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGILL, syscall.SIGTERM)
	fmt.Printf("quit(%v)\n", <-sig)
}
