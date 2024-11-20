package main

import (
	"fmt"

	"github.com/bloiseleo/CodeCrossFit/datastructures/arrays"
)

func main() {
	arr := arrays.CreateDyanmicArray()
	for i := 0; i < 11; i++ {
		arr.Append(i)
	}
	arr.Delete(7)
	fmt.Println("Tamanho", arr.Length())
	arr.Delete(7)
	fmt.Println("Tamanho", arr.Length())
	for i := 0; i < arr.Length(); i++ {
		staticArr := arr.ToStaticArray()
		fmt.Println(staticArr[i])
	}
}
