package euler

import (
	"fmt"
	"math"
	"strconv"
)


func CheckPrime() func(intPrime int64) (bool, *[]int64) {
	primes := []int64{2}
	return func(intPrime int64) (bool, *[]int64) {
		nextPrime := primes[len(primes)-1] + 1
		if intPrime < primes[len(primes)-1] {
			for _, confirmedPrime := range primes {
				if intPrime == confirmedPrime {
					return true, &primes
				}
			}
		}
		isPrime := true
		for nextPrime <= intPrime{
			isPrime = true
			for _, confirmedPrime := range primes {
				if nextPrime%confirmedPrime == 0 {
					isPrime = false
					nextPrime++
					break
				}
			}
			if isPrime {
				primes = append(primes, nextPrime)
				if intPrime == nextPrime {
					return true, &primes
				}
			}
		}
		return isPrime, &primes
	}

}

// calculates the sum of the multiples of a number up to limit
func AddMultiples(numbers []int, limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		for _, number := range numbers {
			if (i % number) == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}

//calculates the sum of even, uneven or all numbers up to limit. set even to nil to sum even and uneven numbers
func FibonacciSum(limit int, even interface{}) int {
	sum := 0
	cache := 1
	for current := 1; current < limit; {
		switch even.(type) {
		case bool:
			if even.(bool) && current%2 == 0 {
				sum += current
			} else if !even.(bool) && current%2 != 0 {
				sum += current
			}
		default:
			sum += current
		}
		cache, current = current, (cache + current)

	}
	return sum
}

//returns biggest prime factor for a number
func BiggestPrimeFactor(number int) int64 {
	rest := int64(number)
	pc := CheckPrime()
	factors := make([]int64, 0)
	limit := int64(math.Sqrt(float64(number)))
	for i := int64(3); i <= limit; i += 2 {
		isPrime, primeList := pc(i)
		if isPrime {
			for _, v := range *primeList {
				if int64(rest)%v == 0 {
					rest = rest / v
					factors = append(factors, v)
					if rest == 1 {
						return v
					}
				}
			}
		}
		if i >= limit-1 {
			return rest
		}
	}
	return 0
}
func BiggestPalindromeProduct(digits int) int64 {
	limit := int64(1)
	for i := 1; i <= digits; i++ {
		limit *= 10
	}
	biggestPalindrome := int64(0)
	for x := limit; x > 0; x-- {
		for y := limit; y > 99; y-- {

			test := x * y
			testString := strconv.FormatInt(test, 10)
			palindrome := true
			nz := len(testString) - 1
			for z := 0; z < len(testString)/2+1; z++ {
				if testString[z] != testString[nz] {
					palindrome = false
					break
				}
				nz--
			}
			if palindrome {
				if x*y > biggestPalindrome {
					biggestPalindrome = x * y
				}

			}
		}
	}
	fmt.Println("returning ", biggestPalindrome)
	return biggestPalindrome
}
func SmallestMultiple(limit int64) int64 {
	for i := limit * (limit - 1); i != 0; i += limit * (limit - 1) {
		noRest := true
		for a := limit - 1; a > limit/2; a-- {
			if i%a != 0 {
				noRest = false
				break
			}
		}
		if noRest {
			return int64(i)
		}
	}
	return 0
}
