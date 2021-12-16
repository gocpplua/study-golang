package main
import "fmt"

func main(){
	var sum = 0
	// Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的
	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Printf("sum is: %v \n", sum)
}