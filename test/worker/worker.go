package worker

import (
	"fmt"
	"time"
)

type Worker struct {
	C    chan int
	Name string
}

type Producer struct {
}

type Consumer struct {
}

func BuildWorker() *Worker {
	return &Worker{
		C:    make(chan int),
		Name: "worker",
	}
}

func (this *Worker) Create() *Worker {
	fmt.Println("Create()")
	return this
}

func (this *Worker) Start() {
	fmt.Println("Start()")

	go func() {
		timeout := time.After(time.Second * time.Duration(1))
		go func() {
			for {
				select {
				case <-timeout:
					{
						fmt.Println("producer", 1)
						this.C <- 1
						close(this.C)
					}
				}
			}
		}()
	}()

	go func() {
		for {
			value, ok := <-this.C
			if ok {
				fmt.Println("consumer", value)
			} else {
				fmt.Println("close")
				break
			}
		}
	}()
}

func (this *Worker) Stop() {
	fmt.Println("Stop()")
}
