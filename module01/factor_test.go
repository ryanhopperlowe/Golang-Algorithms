package module01

import (
	"fmt"
	"sort"
	"testing"
)

func TestFactor(t *testing.T) {
	allPrimes := []int{2}
	var prev float64 = 1
	for i := 0; i < 2000; i++ {
		prime := prev + 2

		isActuallyPrime := true
		for _, p := range allPrimes {
			if int(prime)%int(p) == 0 {
				isActuallyPrime = false
				break
			}
		}
		if isActuallyPrime == true {
			allPrimes = append(allPrimes, int(prime))
		}
		prev = prime
	}

	fmt.Println(allPrimes)

	// tenPrimes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 977, 1667, 1993, 149711}
	tests := []struct {
		primes []int
		number int
		want   []int
	}{
		{allPrimes, 30, []int{2, 3, 5}},
		{allPrimes, 28, []int{2, 2, 7}},
		{allPrimes, 720, []int{2, 2, 2, 2, 3, 3, 5}},
		{allPrimes, 30, []int{2, 3, 5}},
		{allPrimes, 720, []int{2, 2, 2, 2, 3, 3, 5}},
		{allPrimes, 4, []int{2, 2}},
		{[]int{}, 4, []int{4}},
		{allPrimes, 92387432, []int{2, 2, 2, 41, 281669}},
		{allPrimes, 496, []int{2, 1469371469}},
	}
	for _, tc := range tests {
		t.Run(fmt.Sprintf("%v with primes", tc.number), func(t *testing.T) {
			got := Factor(tc.primes, tc.number)
			sort.Ints(got)
			sort.Ints(tc.want)
			if err := intSlicesEqual(got, tc.want); err != nil {
				t.Fatalf("Factor() sorted = %v; want %v; err = %v", got, tc.want, err)
			}
		})
	}
}
