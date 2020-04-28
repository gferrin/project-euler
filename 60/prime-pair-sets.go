package main

import "fmt"

// The idea I have to solve this problem is first to look for all sets of prime pairs
// and it's nessisarily true that all quadrupels are made of tuples and all sets of five
// contain a valid set of 4 so use that idea to build up from tuples to all sets of five
// (below some reasonable number of primes) and sum them all up and see which is min

func main() {
	// the set of primes you're checking against should consist of
	// larger numbers then the one for constructing tuples
	primes := findPrimesBelow(1000)
	bigPrimes := findPrimesBelow(1000000)
	tuples := findTuples(primes, bigPrimes)

	fmt.Println(tuples)
}

func findTuples(primes []int, bigPrimes []int) [][4]int {
	var pairs [][4]int
	fmt.Println("got big primes", len(bigPrimes))
	// first iterate over the primes to find the set of tuples
	for _, px := range primes {
		for _, py := range primes {
			candidatePair := buildPairs(px, py)
			if isPrime(candidatePair[0], bigPrimes) {
				if isPrime(candidatePair[1], bigPrimes) {
					pairs = append(pairs, candidatePair)
				}
			}
		}
	}
	return pairs
}

func isPrime(x int, primes []int) bool {
	for _, prime := range primes {
		if x == prime {
			return true
		}
	}
	return false
}

func buildPairs(x, y int) [4]int {
	pair := [4]int{0, 0, x, y}

	xmultiple := 1
	xc := x
	for xc > 0 {
		xmultiple *= 10
		xc = xc / 10
	}

	ymultiple := 1
	yc := y
	for yc > 0 {
		ymultiple *= 10
		yc = yc / 10
	}

	pair[0] = x*ymultiple + y
	pair[1] = y*xmultiple + x

	return pair
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
