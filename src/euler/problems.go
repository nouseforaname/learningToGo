package euler

import (
	"math"
	"reflect"
)

type Prime int

func (p *Prime) checkPrime(primes []int64) bool {
	isPrime := true
	intPrime := reflect.ValueOf(*p).Int()
	primeTestLimit := math.Sqrt(float64(intPrime))
	for _, confirmedPrime := range primes {
		if intPrime%confirmedPrime == 0 {
			isPrime = false
			break
		} else if float64(confirmedPrime) < primeTestLimit {
			break
		}
	}
	return isPrime
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

//the solution is not optimal but i can play around with interfaces, methods and pointers.. probably a closure with an anonymous
//function will be better for calculating primes. i dont know, im learning. so lets call this the way i currently understand
func BiggestPrimeFactor(number int) int64 {
	rest := int64(number)
	var primes []int64
	primes = make([]int64, 0)
	primes = append(primes, 2)
	factors := make([]int64, 0)
	limit := int64(math.Sqrt(float64(number)))
	for i := int64(3); i <= limit; i += 2 {
		p := Prime(i)
		isPrime := p.checkPrime(primes)
		if isPrime {
			primes = append(primes, int64(i))
			for _, v := range primes {
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
