package module01

import (
	"fmt"
)

func GCD(a, b int) int {
	if b == 0 {
		return a
	}
	for b != 0 {
		fmt.Println(a, "%", b, "=", a%b)
		a, b = b, a%b
	}
	fmt.Println(a)
	return a
}

// func GCD(a, b int) int {
// 	primes := GetNPrimeNumbers(50)

// 	aCounts, fCounts := make(map[int]int), make(map[int]int)
// 	for _, af := range Factor(primes, a) {
// 		aCounts[af]++
// 	}

// 	for _, f := range Factor(primes, b) {
// 		if aCounts[f] > fCounts[f] {
// 			fCounts[f]++
// 		}
// 	}

// 	gcd := 1
// 	for i, c := range fCounts {
// 		gcd *= i * c
// 	}

// 	// fmt.Println(aCounts, fCounts, gcd)
// 	return gcd
// }

func GetNPrimeNumbers(n int) []int {
	primes := []int{2}
	current := 1
	for len(primes) < n {
		current += 2
		isPrime := true
		for _, i := range primes {
			if current%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime == true {
			primes = append(primes, current)
		}
	}
	// fmt.Println(primes)
	return primes
}
