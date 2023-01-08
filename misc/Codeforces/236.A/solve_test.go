package main

import "testing"

func Test_isMale(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"wjmzbmr", false},
		{"xiaodao", true},
		{"sevenkplus", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMale(tt.name); got != tt.want {
				t.Errorf("isMale() = %v, want %v", got, tt.want)
			}
		})
	}
}
