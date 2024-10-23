package main

import "fmt"

// Outputs all prime numbers from 2 to the maximum number inclusive in the Go / Golang\language. 
// Выводит все простые числа от 2 до максимального числа включительно на языке Go/Golang. 

func printPrimes(max int) {
	
	for n := 2; n <= max; n++ {
	  flag := false
		if n == 2 {
			fmt.Println(n)
		}
		if n%2 == 0 {
			continue
		}
		for i := 3; i * i <= n; i += 2 {
			if n % i == 0 {
			 	flag = true
				break
			}
		}
		if flag {
			continue
		}
		fmt.Println(n)
	}

}

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
