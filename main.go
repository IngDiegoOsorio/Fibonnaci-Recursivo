package main
import "fmt"

func fibonacci(limite uint) uint{
	if limite == 0 || limite == 1 {
		return limite
	}
	return fibonacci(limite-2) + fibonacci(limite-1)
}

func imprimir_fibonacci(iteraciones uint){
	var n uint
	for n = 0 ; n < iteraciones; n++ {
		fmt.Printf("| %d ", fibonacci(n))
	  }
	  fmt.Printf("|\n")
}
func main(){
	
	imprimir_fibonacci(15)

}