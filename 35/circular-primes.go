package main

import (
	"fmt"
	"math"
)

func main() {
	circularPrimes := 0
	primes := findPrimesBelow(1000000)
	for _, prime := range primes {
		allPrime := true
		for _, rotation := range rotateNumber(prime) {
			isPrime := false
			for _, p := range primes {
				if rotation == p {
					isPrime = true
					break
				}
			}
			if !isPrime {
				allPrime = false
				break
			}

		}
		if allPrime {
			circularPrimes++
		}
	}
	fmt.Println(circularPrimes)
}

func rotateNumber(n int) []int {
	elements := make([]int, 0)
	for n > 0 {
		elements = append(elements, n%10)
		n = n / 10
	}
	// reverse order of elements
	digits := len(elements)
	orderedElements := make([]int, digits)
	for i, e := range elements {
		orderedElements[digits-i-1] = e
	}
	// create rotations
	rotations := make([]int, digits)
	for i := 0; i < digits; i++ {
		rotation := 0
		for j := 0; j < digits; j++ {
			rotation += orderedElements[(j+i)%digits] * int(math.Pow10(digits-j-1))
		}
		rotations[i] = rotation
	}
	return rotations
}

func findPrimesBelow(n int) []int {
	primes := make([]int, 0)
	primes = append(primes, 2, 3)
	for i := 4; i <= n; i++ {
		isPrime := true
		for _, prime := range primes {
			// if candidate number is devisible by a prime then it's not a prime
			if i%prime == 0 {
				isPrime = false
				break

			} else if prime*prime > i {
				// only need to check numbers below the square root but that function takes a floating point so doing this instead
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}
