package math

func Gcd(a, b int) float64 {
	if b == 0 {
		return float64(a)
	}

	return Gcd(b, a%b)
}
