package codility

func Fibo(a int) int {
	if a == 0 {
		return 0
	}

	if a == 1 {
		return 1
	}

	return Fibo(a-1) + Fibo(a-2)
}
