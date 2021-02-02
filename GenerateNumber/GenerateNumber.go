package main

import (
	"fmt"
	"os"
)

func main() {
	
	N :=1

	for {

		fmt.Print("\nНатуральное число: ")
        fmt.Fscan(os.Stdin, &N)
        if N == 0 {
            break
        }
		fmt.Println("Квадрат числа = ", square(N))
	}

}

func square(N int) int {
	return N * N
}
