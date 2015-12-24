package main

import (
	"fmt"
	"time"
	"errors"
)

func main() {
	fmt.Println("Welcome to the playground!")

	fmt.Println("The time is", time.Now())

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}
	
	var array [10]int
	slice := array[2:4]
	fmt.Println("slice : ", slice, cap(slice), len(slice))//slice :  [0 0] 8 2
	slice = array[2:4:7]
	fmt.Println("slice : ", slice, cap(slice), len(slice))//slice :  [0 0] 5 2

}
