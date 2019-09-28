package main

import "fmt"

func test1() {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()

	panic("触发异常")
}

func main() {
	test1()
}

/*
1、
打印后
打印中
打印前
panic: 触发异常

defer 的执行顺序是后进先出，其实可以看做是存储在一个栈中。
当出现 panic 语句的时候，会先按照 defer 的后进先出的顺序执行，最后才会执行panic
*/
