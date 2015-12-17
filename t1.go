package main

import (
	"fmt"
	"time"
	"os/exec"
)

func main() {
	str := "Welcome to the advertisement! "
	start := 0
	end := 8
	for {
		fmt.Println(str[start:end])
		time.Sleep(200 * time.Millisecond);
		str = str[1:] + str[:1]
		cmd := exec.Command("cls");
		cmd.Run()
	}
}
