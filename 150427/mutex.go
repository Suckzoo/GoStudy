package main

import "fmt"
import "runtime"
import "sync"

func main() {
	runtime.GOMAXPROCS(10)
	sum := 0
	rout1 := make(chan int)
	rout2 := make(chan int)
	var mutex = &sync.Mutex{}
	go func() {
		for i:=1;i<=1000000;i++ {
			mutex.Lock()
			sum++
			mutex.Unlock()
		}
		rout1<-1
	}()
	go func() {
		for i:=1;i<=1000000;i++ {
			mutex.Lock()
			sum++
			mutex.Unlock()
		}
		rout2<-1
	}()
	<-rout1
	<-rout2
	fmt.Println(sum)
}
