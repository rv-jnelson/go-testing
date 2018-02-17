package primes_test

import (
	"testing"
	"github.com/rv-jnelson/go-testing/benchmark"
)

//var primeValse = [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103, 107, 109, 113, 127, 131, 137, 139, 149, 151, 157, 163, 167, 173, 179, 181, 191, 193, 197, 199}

func TestLargestPrimeUnder(t *testing.T) {
	p := primes.LargestPrimeUpTo(28)
	if p != 23 {
		t.Error("prime calc is wrong, expected 28 but got", p)
	}


	p = primes.LargestPrimeUpTo(29)
	if p != 29 {
		t.Error("prime calc is wrong, expected 29 but got", p)
	}

	p = primes.LargestPrimeUpTo(4)
	if p != 3 {
		t.Error("prime calc is wrong, expected 3 but got", p)
	}
}


func benchPrime(x int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		primes.LargestPrimeUpTo(x)
	}
}

func BenchmarkLargestPrimeUpTo6(b *testing.B) {benchPrime(6, b)}
func BenchmarkLargestPrimeUpTo36(b *testing.B) {benchPrime(36, b)}
func BenchmarkLargestPrimeUpT72(b *testing.B) {benchPrime(72, b)}
func BenchmarkLargestPrimeUpTo144(b *testing.B) {benchPrime(144, b)}
func BenchmarkLargestPrimeUpTo288(b *testing.B) {benchPrime(288, b)}
func BenchmarkLargestPrimeUpTo376(b *testing.B) {benchPrime(376, b)}
func BenchmarkLargestPrimeUpTo1024(b *testing.B) {benchPrime(1024, b)}
func BenchmarkLargestPrimeUpTo2048(b *testing.B) {benchPrime(2048, b)}
func BenchmarkLargestPrimeUpTo4096(b *testing.B) {benchPrime(4096, b)}
func BenchmarkLargestPrimeUpTo8192(b *testing.B) {benchPrime(8192, b)}
func BenchmarkLargestPrimeUpTo16384(b *testing.B) {benchPrime(16384, b)}