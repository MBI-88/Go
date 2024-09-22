package benchmark

import (
	"testing"
)


/*
	Benchmarking is the practice of measuring the performance of a program on a fixed
	workload. In Go, a benchmark function looks like a test function, but with the Benchmark
	prefix and a *testing.B parameter that provides most of the same methods as a *testing.T,
	plus a few extra related to performance measurement. It also exposes an integer field N, which
	specifies the number of times to perform the operation being measured.

	The report tells us that each call to IsPalindrome to ok about 1.035 microseconds, averaged
	over 1,000,000 runs. Since the benchmark runner initially has no idea how long the operation
	takes, it makes some initial measurements using small values of N and then extrapolates to a
	value large enough for a stable timing measurement to be made.

*/

func BenchmarkIsPlindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsPalindrome("A man, a plan a canal: Panama")
	}
}


