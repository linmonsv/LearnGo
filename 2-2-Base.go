package main

import (
	"errors"
	"fmt"
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
}
