/*
 * PandaOS
 * Copyright (C) 2025  Cyberpandino
 */

package server

import "testing"

func TestItoa(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{0, "0"},
		{1, "1"},
		{10, "10"},
		{42, "42"},
		{100, "100"},
		{-1, "-1"},
		{-42, "-42"},
		{999999, "999999"},
		{1000000, "1000000"},
	}
	for _, c := range cases {
		got := itoa(c.in)
		if got != c.want {
			t.Errorf("itoa(%d) = %q, want %q", c.in, got, c.want)
		}
	}
}

func BenchmarkItoa(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		itoa(123456)
	}
}
