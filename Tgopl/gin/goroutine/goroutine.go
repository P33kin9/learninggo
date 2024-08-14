package main

import (
	"fmt"
	"time"
)

func worker(id int) {
	fmt.Println("Worker", id, "started")
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "finished")

}

func main() {
	for i := 0; i < 5; i++ {

		go worker(i)
	}

	time.Sleep(time.Second * 5)

}
