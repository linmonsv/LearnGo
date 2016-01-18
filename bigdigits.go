package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var bigDigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{
		"    1  ",
		"   11  ",
		"    1  ",
		"    1  ",
		"    1  ",
		"    1  ",
		"   111 ",
	},
	{
		"   222 ",
		"  2   2",
		"     2 ",
		"    2  ",
		"   2   ",
		"  2    ",
		" 22222 ",
	},
	{
		"   333 ",
		" 3    3",
		"      3",
		"   333 ",
		"      3",
		" 3    3",
		"   333 ",
	},
	{
		"    4  ",
		"   44  ",
		"  4 4  ",
		" 4  4  ",
		" 444444",
		"    4  ",
		"    4  ",
	},
	{
		" 55555 ",
		" 5     ",
		" 5     ",
		" 5555  ",
		"     5 ",
		" 5   5 ",
		"  555  ",
	},
	{
		"   66  ",
		"  6    ",
		" 6     ",
		" 6666  ",
		" 6   6 ",
		" 6   6 ",
		"  666  ",
	},
	{
		" 77777 ",
		"      7",
		"    7  ",
		"   7   ",
		"  7    ",
		"  7    ",
		" 7     ",
	},
	{
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
		" 8   8 ",
		" 8   8 ",
		"  888  ",
	},
	{
		"  9999 ",
		" 9   9 ",
		" 9   9 ",
		"  9999 ",
		"     9 ",
		"     9 ",
		"     9 ",
	},
}

func getRandomNums() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	stringOfDigits := ""
	for i := 0; i < 8; i++ {
		stringOfDigits += fmt.Sprintf("%d", r.Intn(10))
	}
	return stringOfDigits
}

func showDigits() {
	//stringOfDigits := "0123456789" //getRandomNums()
	stringOfDigits := getRandomNums()
	//stringOfDigits = "57235742"
	//fmt.Println(time.Now().Format("20060102150405"))
	stringOfDigits = time.Now().Format("20060102150405")
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			}
		}
		fmt.Println(line)
	}

}

func main() {
	for {
		showDigits()
		time.Sleep(1000 * time.Millisecond)

		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()

		//break
	}
}
