package main

import "os"
import "fmt"

func main() {
	argsProg := os.Args
	argWithoutProg := os.Args[1:]
	arg := os.Args[3]
	fmt.Println(argsProg)
	fmt.Println(argWithoutProg)
	fmt.Println(arg)
}
