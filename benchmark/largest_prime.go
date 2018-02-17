package primes

func LargestPrimeUpTo(x int) int {
	for i:= x; i >= 3; i = i - 1 {
		isPrime := true
		for j := 2; j < i; j = j + 1 {
			if i % j == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			return i
		}
	}

	return 3
}
