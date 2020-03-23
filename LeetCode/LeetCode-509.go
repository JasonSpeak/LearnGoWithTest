package main

func fib(N int) int {
	preNums := [2]int{0, 1}
	if N < 2 {
		return preNums[N]
	}

	for i := 2; i < N; i++ {
		nextNum := preNums[0] + preNums[1]
		preNums[0] = preNums[1]
		preNums[1] = nextNum
	}

	return preNums[1]
}
