package main

import (
	"fmt"
)

func main() {

	fmt.Println("------ 流程控制 ------")
	fmt.Println("------ ---if ------")
	x := 10
	//Go里面if条件判断语句中不需要括号
	if x > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	//Go的if还有一个强大的地方就是条件判断语句里面允许声明一个变量，
	//这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	if x1 := 12; x1 > 10 {
		fmt.Println("x is greater than 10")
	} else {
		fmt.Println("x is less than 10")
	}

	integer := 4
	if integer == 3 {
		fmt.Println("The integer is equal to 3")
	} else if integer < 3 {
		fmt.Println("The integer is less than 3")
	} else {
		fmt.Println("The integer is greater than 3")
	}

	fmt.Println("------ ---goto ------")
	x2 := 0
Here: //标签名是大小写敏感的。
	x2++
	if x2 < 5 {
		goto Here
	}
	fmt.Println("x2 :", x2)

	fmt.Println("------ ---for ------")
	//Go里面最强大的一个控制逻辑就是for，
	//它即可以用来循环读取数据，
	//又可以当作while来控制逻辑，
	//还能迭代操作
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to", sum)

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum is equal to", sum)

	for index := 10; index > 0; index-- {
		if index == 5 {
			break //break or continue
		}
		fmt.Println(index)
	}

	numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	for k, v := range numbers {
		fmt.Println(k, v)
		fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

	//由于 Go 支持 “多值返回”, 而对于“声明而未被调用”的变量, 编译器会报错,
	//在这种情况下, 可以使用_来丢弃不需要的返回值
	for _, v := range numbers {
		//fmt.Println(k, v)
		//fmt.Println("map's key:", k)
		fmt.Println("map's val:", v)
	}

	fmt.Println("------ ---switch ------")
	//Go的switch非常灵活，表达式不必是常量或整数，
	//执行的过程从上至下，直到找到匹配项；而如果switch没有表达式，它会匹配true
	i := 10
	switch i {
	//Go里面switch默认相当于每个case最后带有break，
	//匹配成功后不会自动向下执行其他case，而是跳出整个switch
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}

	//但是可以使用fallthrough强制执行后面的case代码
	integer = 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")

	}

	fmt.Println("------ 函数 ------")
}
