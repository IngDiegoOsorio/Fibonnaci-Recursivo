package main

import "fmt"

func sum(args... int) int{
	mayor := 0
	for _, v := range args {
		if(v > mayor){
			mayor = v
		}
	}
	return mayor
}

func main() {
	numeros := [] int {5,20,6,14}
	fmt.Println(sum(numeros...))
}