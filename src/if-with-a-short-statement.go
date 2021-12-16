package main
import "fmt"

func main(){
	if sum := 1; sum <= 0 {  // var sum = 0 -> error:  var declaration not allowed in if initializer
		fmt.Println("less than 0")
	}	else {
		fmt.Println(sum)
	}
}