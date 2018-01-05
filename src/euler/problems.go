package euler

import (
	"fmt"
	"math"
	"strconv"
)


func CheckPrime() func(intPrime int64) (bool, *[]int64) {
	primes := []int64{2}
	return func(intPrime int64) (bool, *[]int64) {
		possiblePrime := primes[len(primes)-1] + 1
		if intPrime < primes[len(primes)-1] {
			for _, confirmedPrime := range primes {
				if intPrime == confirmedPrime {
					return true, &primes
				}
			}
		}
		isPrime := true
		for possiblePrime <= intPrime{
			isPrime = true
			for _, confirmedPrime := range primes {
				if possiblePrime%confirmedPrime == 0 {
					isPrime = false
					possiblePrime++
					break
				}
			}
			if isPrime {
				primes = append(primes, possiblePrime)
				if intPrime == possiblePrime {
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

func PrimeFactor(number int) []int64 {
	rest := int64(number)
	fmt.Println("getting Primefactorization for ",number)
	pc := CheckPrime()
	factors := make([]int64,0)	
	
	limit:=int64(number)
	if limit>100{
		limit = int64(math.Sqrt(float64(number)))
	}
	for i := int64(1); i <= limit; i += 2 {
		isPrime, primeList := pc(i)
		if isPrime {
			for _, v := range *primeList {
				if int64(rest)%v == 0 {
					rest = rest / v
					factors = append(factors, v)
					if rest == 1 {
						return factors
					}
				}
			}
		}
	}
	return factors
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
func SmallestMultipleV2(limit int) int64 {
	factorCollection := make(map[int64]int64)
	for i:=2; i<=limit; i++{
		factorsI:=PrimeFactor(i)
		factorMapI := make(map[int64]int64)
		
		for _,v := range factorsI {
			_, exists := factorMapI[v]
			if exists {
				factorMapI[v]+=1
			}else{
				factorMapI[v]=1
			}
		}
		for k , v := range factorMapI{
			_, exists := factorCollection[k]
			if exists {
				if factorCollection[k]<factorMapI[k]{
					factorCollection[k]=factorMapI[k]
				}
			}else{
				factorCollection[k]=v
			}			
		}
	}
	sum:=int64(1)
	for k,v	:= range factorCollection{
		for i:=int64(1); i<=v;i++{
			sum*=k
		}
	}
	return sum
}