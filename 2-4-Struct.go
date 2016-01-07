package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

//struct的匿名字段
//当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct
type Student struct {
	person
	score int
	name  string
}

func main() {
	var p person
	p.name = "Astaxie"
	p.age = 25
	fmt.Printf("The person's name is %s\n", p.name)
	p1 := person{"Tom", 27}
	p2 := person{age: 24, name: "Cat"}
	fmt.Printf("%v %v %v\n", p, p1, p2)

	pp, pp_diff := Older(p1, p2)
	fmt.Printf("%v %d\n", pp, pp_diff)

	//匿名字段就是这样，能够实现字段的继承
	//重载通过匿名字段继承的一些字段
	//最外层的优先访问
	stu := Student{person{"Mark", 18}, 99, "XiaoMark"}
	fmt.Printf("%v\n", stu)
	fmt.Printf("%s %d %d %s\n", stu.name, stu.age, stu.score, stu.person.name)

}

func Older(p1, p2 person) (person, int) {
	if p1.age > p2.age {
		return p1, p1.age - p2.age
	}
	return p2, p2.age - p2.age
}
