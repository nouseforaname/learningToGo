package euler

import (

)

// calculates the sum of the multiples of a number up to limit
func AddMultiples(numbers []int, limit int) int{
	sum:=0;
	for i:=1; i<=limit; i++ {
		for _,number:=range numbers{
			if (i%number) == 0 {
				sum+=i
				break
			}
		}
	}
	return sum
}