package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	str := "Welcome to the advertisement! "
	start := 0
	end := 8
	for {
		fmt.Println(str[start:end])
		time.Sleep(200 * time.Millisecond)
		str = str[1:] + str[:1]
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
