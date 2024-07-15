package fibonacci

import "fmt"

func View() {
	fmt.Println("Enter sequence element number: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("Fibonacci sequence number:", fibonacci(n))
}

func fibonacci(n int) uint64 {
	var f []uint64 = []uint64{0, 1}
	for len(f) <= n {
		f = append(f, f[len(f)-1]+f[len(f)-2])

	}
	fmt.Println("The fibonacci sequence: ")
	fmt.Println(f)
	return f[n]
}
