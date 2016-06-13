package main

import "time"
import "fmt"

func main() {
	burstyLimiter := make(chan time.Time, 3)
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			fmt.Println("Tick: ",t)
			burstyLimiter <- t
		}
	}()
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for i := range burstyRequests {
		<- burstyLimiter
		fmt.Println("Job", i, "done: ", time.Now())
	}
}
