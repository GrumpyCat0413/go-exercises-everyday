package main

import "fmt"

func test10() {
	/*
		通过指针变量 p 访问其成员变量 name，有哪几种方式？ AC
		A.p.name
		B.(&p).name
		C.(*p).name
		D.p->name
	*/
}

func test11() {
	type MyInt1 int   // redefine type
	type MyInt2 = int // type alias
	var i int = 0
	var i1 MyInt1 = i //cannot use i (type int) as type MyInt1 in assignment
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func main() {
	test10()
	test11()
}

/*
10、
AC。& 取址运算符，* 指针解引用。
*/

/*
11、
类型别名与类型定义的区别
type MyInt1 int 基于类型 int 创建了新类型 MyInt1，
type MyInt2 = int 创建了 int 的类型别名 MyInt2

*/
