package main

import "fmt"
import "os/exec"

func main() {
	var command string
	for {
		fmt.Scanln(&command)
		//exit condition?
		if command == "exit" {
			break
		}
		//exec'ing process with exec.Command
		cmd := exec.Command("bash", "-c", command)
		//get output
		cmdOut, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(cmdOut))
	}
}
