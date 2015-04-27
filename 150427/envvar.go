package main

import (
	"os"
	"strings"
	"fmt"
)

func main() {
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()
	for t, e := range os.Environ() {
		fmt.Println(t)
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}
}
