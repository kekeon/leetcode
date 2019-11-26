package array

func fib(N int) int {

	list := []int{0, 1}

	if N < 2 {
		return list[N]
	}

	for i := 2; i <= N; i++ {
		list = append(list, list[i-1]+list[i-2])
	}

	return list[N]
}

func recursiveFib(N int) int {

	if N == 0 {
		return 0
	} else if N <= 2 {
		return 1
	}

	return recursiveFib(N-1) + recursiveFib(N-2)
}

func fib2(N int) int {


	if N <= 0 {
		return 0
	}
	s := 1
	c := 0
	for N--; 0 < N; N-- {
		s = s + c
		c = s - c
	}

	return  s
}
