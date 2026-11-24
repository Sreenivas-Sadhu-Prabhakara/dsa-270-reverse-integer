package main

import "fmt"

const (
	maxInt32 = 1<<31 - 1
	minInt32 = -1 << 31
)

// ReverseInt reverses a signed 32-bit integer's decimal digits, returning 0
// when the result would overflow int32. The overflow test runs BEFORE each
// append, predicting rather than detecting.
func ReverseInt(x int) int {
	out := 0
	for x != 0 {
		digit := x % 10 // Go keeps the sign on the remainder: -123%10 = -3
		x /= 10
		if out > maxInt32/10 || (out == maxInt32/10 && digit > maxInt32%10) {
			return 0
		}
		if out < minInt32/10 || (out == minInt32/10 && digit < minInt32%10) {
			return 0
		}
		out = out*10 + digit
	}
	return out
}

func main() {
	fmt.Println(ReverseInt(123), ReverseInt(-123), ReverseInt(1534236469))
}
