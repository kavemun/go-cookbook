package main

import "testing"

func TestCountTheWays_String(t *testing.T) {
	var c = CountTheWays([]int{1, 2})
	tests := []struct {
		name string
		c    *CountTheWays
		want string
	}{
		{"base-case", &c, "1...2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("CountTheWays.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
