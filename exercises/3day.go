package main

import "fmt"

func test3_1() {
	s := make([]int, 5) //初始化容量为5的切片 每个元素是0，而不是容量为5，没有元素使用
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}

func test3_2() {
	s := make([]int, 0) //初始化容量为0切片1
	fmt.Println(len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))
}

func test4(x, y int) (sum int, error) {
	return x + y, nil
}

func test5() {
	//new() and make()的区别

}

func main() {
	test3_1()
	test3_2()
	test4()
	test5()
}

/*
3、
output:
[0 0 0 0 0 1 2 3]
[1 2 3 4]
*/

/*
4、
第二个返回值没有命名
在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。
如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()。
这里的第一个返回值有命名 sum，第二个没有命名，所以错误
*/

/*
5、
new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同
new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T 的值。
换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。
适用于值类型，如数组、结构体等

make(T,args) 返回初始化之后的 T 类型的值，
这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。
make() 只适用于 slice、map 和 channel
*/
