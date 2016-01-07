package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现Sayhi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	Tom := Employee{Human{"Sam", 36, "444-222-XXX"}, "Things Ltd.", 5000}

	//定义Men类型的变量i
	var i Men

	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")
	//i也能存储Employee
	i = Tom
	fmt.Println("This is Tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")

	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//T这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
	//你会发现interface就是一组抽象方法的集合，
	//它必须由其他非interface类型实现，而不能自我实现，
	//go 通过interface实现了duck-typing

	// 定义a为空接口
	var a interface{}
	var i2 int = 5
	s := "Hello world"
	// a可以存储任意类型的数值
	a = i2
	fmt.Println(a)
	a = s
	fmt.Println(a)

	/*
		fmt.Println是我们常用的一个函数，
		但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文
		件，你会看到这样一个定义:
		type Stringer interface {
			String() string
		}
		也就是说，任何实现了String方法的类型都能作为参数被fmt.Println调用
	*/
	fmt.Println("This Human is :", Tom)

	//反向知道这个变量里面实际保存了的是哪个类型的对象
	//Comma-ok断言
	if _, ok := a.(int); ok {
		fmt.Println("a is int")
	} else if _, ok := a.(string); ok {
		fmt.Println("a is string")
	}

	//element.(type)语法不能在switch外的任何逻辑里面使用，
	//如果你要在switch	外面判断一个类型就使用comma - ok
	a = i2
	switch value := a.(type) {
	case int:
		fmt.Println("a is int ", value)
	case string:
		fmt.Println("a is string ", value)
	}

	//嵌入interface
	//如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method
	/*
		一个例子就是io包下面的 io.ReadWriter ，他包含了io包下面的Reader和Writer两个interface。
		// io.ReadWriter
		type ReadWriter interface {
			Reader
			Writer
		}
	*/

	//reflect
	var x2 float64 = 3.4
	v := reflect.ValueOf(x2)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

	//反射的字段必须是可修改的
	v2 := reflect.ValueOf(&x2)
	v3 := v2.Elem()
	v3.SetFloat(7.1)
	fmt.Println("value:", v3)
}

//需要某个类型能被fmt包以特殊的格式输出，你就必须实现Stringer这个接口。
//如果没有实现这个接口，fmt将以默认的方式输出
func (h Human) String() string {
	return " " + h.name + " - " + strconv.Itoa(h.age) + " years - " + h.phone + " "
}

//实现了error接口的对象（即实现了Error() string的对象），
//使用fmt输出时，会调用Error()方法，因此不必再定义String()方法了
