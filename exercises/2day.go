package main

import "fmt"

func test2() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
		//fmt.Println(&val)
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func main() {
	test2()
}

/*
2、
output:
0 -> 3
1 -> 3
2 -> 3
3 -> 3
for range 循环的时候会创建每个元素的副本，而不是元素的引用，
所以 m[key] = &val 取的都是变量 val 的地址，
所以最后 map 中的所有元素的值都是变量 val 的地址，
因为最后 val 被赋值为3，所有输出都是3.
所以正确写法
value := val
m[key] = &value
*/
