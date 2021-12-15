package main
import "fmt"

func add(x int, y int) int{
	return x + y
}


func addextra(x, y, z int) int {
	return x + y + z
}

func main(){
	fmt.Println(add(1, 3))
	fmt.Println(addextra(1, 3, 5))
	
}