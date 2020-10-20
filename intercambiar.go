package main

import "fmt"

func intercambiar(a, b *int) {
	aux := *a
	*a = *b
	*b = aux

}

func main(){
	a := 2
	b := 4
	fmt.Println(a,b)
	intercambiar(&a,&b)
	fmt.Println(a,b)
}