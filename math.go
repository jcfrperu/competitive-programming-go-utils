package cp

func GCD[T IntegerNumber](a, b T) T {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T IntegerNumber](a, b T) T {
	return a / GCD(a, b) * b
}
