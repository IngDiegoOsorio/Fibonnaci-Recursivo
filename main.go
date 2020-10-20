package main
import "fmt"

func fibonacci(limite uint) uint{
	if limite == 0 || limite == 1 {
		return limite
	}

	return fibonacci(limite-2) + fibonacci(limite-1)
}

func main(){
	var x = fibonacci(10)
	fmt.Println(x)
}