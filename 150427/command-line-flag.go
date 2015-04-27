package main

import "flag"
import "fmt"

func main() {
	wordPtr := flag.String("word", "foo", "a string")
	var floatVal = new(float64)
	flag.Float64Var(floatVal, "gpa", 1, "your gpa value")
	flag.Parse()
	fmt.Println(*wordPtr, *floatVal)
}
