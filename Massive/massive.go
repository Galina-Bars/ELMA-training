package main

import (
	"fmt"
	
)

func main() {
	var K int = 3
	A := [5] int {1, 2, 3, 4, 5}
	fmt.Println("Исходный массив: ")
	fmt.Println(A)
		
	A1 := append(A[K:], A[:K]...)
	fmt.Println("\nМассив после сдвига. Индекс сдвига К =", K)
	fmt.Println (A1)

}		
