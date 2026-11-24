package main

import "testing"

func TestReverseInt(t *testing.T) {
	cases := []struct{ in, want int }{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{1534236469, 0},          // reverses past int32 max
		{-2147483648, 0},         // min int32 reverses way out of range
		{2147483647, 0},          // 7463847412 overflows
		{1463847412, 2147483641}, // just inside the limit
		{7, 7},
	}
	for _, c := range cases {
		if got := ReverseInt(c.in); got != c.want {
			t.Errorf("ReverseInt(%d) = %d, want %d", c.in, got, c.want)
		}
	}
}
