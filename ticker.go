package main

import "time"
import "fmt"
func main() {
	ticker1 := time.NewTicker(time.Millisecond * 500)
	ticker2 := time.NewTicker(time.Millisecond * 1000)
	ticker3 := time.NewTicker(time.Millisecond * 3000)
	go func() {
		for {
			<-ticker1.C
			fmt.Println("띠")
		}
	}()
	go func() {
		for {
			<-ticker2.C
			fmt.Println("용")
		}
	}()
	go func() {
		for t := range ticker3.C {
			fmt.Println(t)
		}
	}()
	time.Sleep(time.Millisecond * 5000)
	ticker1.Stop()
	ticker2.Stop()
	ticker3.Stop()
	fmt.Println("Ticker stopped")
}
