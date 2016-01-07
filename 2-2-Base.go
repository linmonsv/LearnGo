package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	fmt.Println("------ 定义变量 ------")
	var i int = 5
	j := 4

	fmt.Printf("%d %d\n", i, j)

	fmt.Println("------ 常量 ------")

	const Pi float32 = 3.1415926
	const MaxThread = 10

	fmt.Printf("%f %d\n", Pi, MaxThread)

	fmt.Println("------ 内置基础类型 ------")
	fmt.Println("------ ---Boolean ------")
	var isActive bool
	var enabled, disabled = true, false
	valid := false
	fmt.Printf("%v %v %v %v\n", isActive, enabled, disabled, valid)

	fmt.Println("------ ---数值类型 ------")
	var a int8
	var b int32
	//c := a + b mismatched types int8 and int32
	fmt.Printf("%v %v\n", a, b)

	fmt.Println("------ ---复数 ------")
	var c complex64 = 5 + 5i
	fmt.Printf("Value is :　%v\n", c)

	fmt.Println("------ 字符串 ------")
	var frenchHello string
	var emptyString string = ""

	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Ohaiou"
	frenchHello = "Bonjour"

	fmt.Printf("%v %v %v\n", no, yes, maybe)
	fmt.Printf("-%v- %v %v\n", emptyString, japaneseHello, frenchHello)

	//在Go中字符串是不可变的，例如下面的代码编译时会报错
	var s string = "hello"
	//s[0] = 'c'

	//但如果真的想要修改怎么办呢？下面的代码可以实现：
	fmt.Println("------ ---修改 ------")
	sbyte := []byte(s)
	sbyte[0] = 'c'
	s2 := string(sbyte)
	fmt.Printf("%s\n", s2)

	//可以使用+操作符来连接两个字符串
	fmt.Println("------ ---连接+ ------")
	s1 := "hello,"
	m2 := " world"
	a3 := s1 + m2
	fmt.Printf("%s\n", a3)

	//字符串虽不能更改，但可进行切片操作
	fmt.Println("------ ---切片 ------")
	s3 := "c" + s1[1:]
	fmt.Printf("%s\n", s3)

	//如果要声明一个多行的字符串怎么办？可以通过`来声明：
	fmt.Println("------ ---多行 ------")
	s4 := `hello
    world`
	fmt.Printf("%s\n", s4)

	fmt.Println("------ 错误类型 ------")

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Printf("%v\n", err)
	}

	fmt.Println("------ Go数据底层的存储 ------")
	fmt.Println("------ 一些技巧 ------")
	fmt.Println("------ ---分组声明 ------")

	/*const (
	      i = 100
	      pi = 3.1415
	      prefix = "Go_"
	  )
	  var (
	      i func init() {
	          pi float32
	          prefix string
	      }
	  )*/
	fmt.Println("------ ---iota枚举 ------")
	//Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，它默认开始值是0，每调用一次加1：
	const (
		x = iota
		y = iota
		z = iota
		w //常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3
	)
	const v = iota // 每遇到一个const关键字，iota就会重置，此时v == 0

	fmt.Println("------ Go程序设计的一些规则 ------")
	fmt.Println("------ 1.大写字母开头的变量是可导出的，也就是其它包可以读取的，是公用变量;------")
	fmt.Println("------ 2.小写字母开头的就是不可导出的，是私有变量;------")
	fmt.Println("------ 3.大写字母开头的函数也是一样，相当于class中的带public关键词的公有函数;------")
	fmt.Println("------ 4.小写字母开头的就是有private关键词的私有函数;------")

	fmt.Println("------ ---array ------")
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])

	arr1 := [3]int{1, 2, 3}
	arr2 := [10]int{1, 2, 3}  //其中前三个元素初始化为1、2、3，其它默认为0
	arr3 := [...]int{4, 5, 6} //可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArry := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Printf("arr1 : %v\n", arr1)
	fmt.Printf("arr2 : %v\n", arr2)
	fmt.Printf("arr3 : %v\n", arr3)
	fmt.Printf("doubleArray : %v\n", doubleArray)
	fmt.Printf("easyArry : %v\n", easyArry)

	fmt.Println("------ ---slice ------")
	//slice并不是真正意义上的动态数组，而是一个引用类型。slice总是指向一个底层array，
	//slice的声明也可以像array一样，只是不需要长度
	//var fslice []int
	slice := []byte{'a', 'b', 'c', 'd'}
	fmt.Printf("slice : %s\n", slice)

	var arr4 = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a4, b4 []byte
	a4 = arr4[2:5]
	b4 = arr4[3:5]
	fmt.Printf("a4 : %s\n", a4)
	fmt.Printf("b4 : %s\n", b4)

	a4 = arr4[:3]
	fmt.Printf("a4 : %s\n", a4)
	a4 = arr4[5:]
	fmt.Printf("a4 : %s\n", a4)
	a4 = arr4[:]
	fmt.Printf("a4 : %s\n", a4)

	fmt.Printf("len : %d\n", len(a4))
	fmt.Printf("cap : %d\n", cap(a4))
	b4 = append(a4, 'k')
	fmt.Printf("append : %s\n", b4)
	b4 = arr4[3:6]
	//copy(b4, a4)
	copy(b4, a4)
	fmt.Printf("a4 : %s\n", a4) //abcabcghij

	//声明一个key是字符串，值为int的字典,这种方式的声明需要在使用之前使用make初始化
	var numbers map[string]int
	numbers = make(map[string]int)
	//另一种map的声明方式
	//numbers := make(map[string]int)
	numbers["one"] = 1  //赋值
	numbers["ten"] = 10 //赋值
	numbers["three"] = 3
	fmt.Println("第三个数字是: ", numbers["three"])
	numbers["one"] = 11
	fmt.Println("第1个数字是: ", numbers["one"])

	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	//map内置有判断是否存在key的方式
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	cRating, ok := rating["C"]
	if ok {
		fmt.Println("C is in the map and its rating is ", cRating)
	} else {
		fmt.Println("We have no rating associated with C in the map")
	}
	delete(rating, "C") //删除key为C的元素
	cRating2, ok2 := rating["C"]
	if ok2 {
		fmt.Println("C is in the map and its rating is ", cRating2)
	} else {
		fmt.Println("We have no rating associated with C in the map")
	}

	m := make(map[string]string)
	m["Hello"] = "Bonjour"
	m1 := m
	m1["Hello"] = "Balut"
	fmt.Println(m["Hello"]) //现在m["hello"]的值已经是Salut了

	fmt.Println("------ ---make、new操作 ------")
	//make用于内建类型（map、slice 和channel）的内存分配。
	//new用于各种类型的内存分配。

	fmt.Println("------ 函数 ------")

	x2 := 3
	y2 := 4
	max_xy := max(x2, y2)
	fmt.Printf("max_xy : %d\n", max_xy)

	fmt.Println("------ ---多个返回值 ------")

	xaddy, xplusy := SumAndProduce(x2, y2)
	fmt.Printf("xaddy, xplusy : %d %d\n", xaddy, xplusy)

	xaddy2, xplusy2 := SumAndProduce2(x2, y2)
	fmt.Printf("xaddy2, xplusy2 : %d %d\n", xaddy2, xplusy2)

	fmt.Println("------ ---变参 ------")
	myFunc(1, 2, 3, 4, 5)

	fmt.Println("------ ---传值与传指针 ------")

	a5 := 3
	a6 := add1(a5)
	fmt.Printf("a5, a6 : %d %d\n", a5, a6)

	a7 := 3
	a8 := add2(&a7)
	fmt.Printf("a7, a8 : %d %d\n", a7, a8)

	fmt.Println("------ ---defer ------")
	//可以在函数中添加多个defer语句。当函数执行到最后时，
	//这些defer语句会按照逆序执行，最后该函数返回
	defer fmt.Println("main end")
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i) //4 3 2 1 0
	}
	fmt.Println("")

	fmt.Println("------ ---数作为值、类型 ------")
	//在Go中函数也是一种变量，我们可以通过type来定义它，
	//它的类型就是所有拥有相同的参数，相同的返回值的一种类型
	slice2 := []int{1, 2, 3, 4, 5, 7}
	fmt.Println("slice2 =", slice2)
	odd := filter(slice2, isOdd)
	fmt.Println("odd =", odd)
	even := filter(slice2, isEven)
	fmt.Println("even =", even)

	fmt.Println("------ ---Panic和Recover ------")
	//Go没有像Java那样的异常机制，它不能抛出异常，而是使用了panic和recover机制
	var user = os.Getenv("USER")
	if user == "" {
		panic("no value for $USER")
	}

	fmt.Println("------ ---main函数和init函数 ------")
	//Go里面有两个保留的函数：
	//init函数（能够应用于所有的package）和main函数（只能应用于package main）。
	//这两个函数在定义时不能有任何的参数和返回值。
	//Go程序会自动调用init()和main()，所以你不需要在任何地方调用这两个函数
	//每个package中的init函数都是可选的，
	//但package main就必须包含一个main函数。

	//main import 包 import 关联包
	//const
	//var
	//init()
	//main()
	//Exit

	fmt.Println("------ ---import ------")
	//fmt是Go语言的标准库，其实是去goroot下去加载该模块
	//Go的import还支持如下两种方式来加载自己写的模块
	//1. 相对路径import “./model” //当前文件同一目录的model目录，但是不建议这种方式来import
	//2. 绝对路径import “shorturl/model” //加载gopath/src/shorturl/model模块

	//一些特殊的import
	/*
		1. 点操作,省略前缀的包名
		import(
			. "fmt"
		)
		fmt.Println("hello world")可以省略的写成Println("hello world")

		2. 别名操作
		别名操作顾名思义我们可以把包命名成另一个我们用起来容易记忆的名字
		import(
			f "fmt"
		)
		别名操作的话调用包函数时前缀变成了我们的前缀，即f.Println("hello world")

		3. _操作
		这个操作经常是让很多人费解的一个操作符，请看下面这个import
		import (
			"database/sql"
			_ "github.com/ziutek/mymysql/godrv"
		) _
		操作其实是引入该包，而不直接使用包里面的函数，而是调用了该包里面的init函数。
	*/
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func SumAndProduce(a, b int) (int, int) {
	return a + b, a * b
}

//最好命名返回值，因为不命名返回值，虽然使得代码更加简洁了
func SumAndProduce2(a, b int) (add int, multiplied int) {
	add = a + b
	multiplied = a * b
	return
}

//arg ...int告诉Go这个函数接受不定数量的参数。
//注意，这些参数的类型全部是int。在函数体中，变量arg是一个int的slice
func myFunc(arg ...int) {
	for _, n := range arg {
		fmt.Printf("And the number is: %d\n", n)
	}
}

func add1(a int) int {
	a = a + 1
	return a
}

//Go语言中string，slice，map这三种类型的实现机制类似指针，
//所以可以直接传递，而不用取地址后传递指针。
//（注：若函数需改变slice的长度，则仍需要取地址传递指针）
func add2(a *int) int {
	*a = *a + 1
	return *a
}

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}

//这个函数检查作为其参数的函数在执行时是否会产生panic
//没明白,,,这例子也,,,
func throwsPanic(f func()) (b bool) {
	defer func() {
		if x := recover(); x != nil {
			b = true
		}
	}()
	f() //执行函数f，如果f中出现了panic，那么就可以恢复回来
	return
}
