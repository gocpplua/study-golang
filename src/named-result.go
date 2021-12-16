package main

import "fmt"

// Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。
// 没有参数的 return 语句返回已命名的返回值
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17)) // output: 7 10
}
