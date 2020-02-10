package stringutil

import "testing"

func TestReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		args string
		want string
	}{
		// TODO: Add test cases.
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}
	for _, tt := range tests {
		t.Run("-", func(t *testing.T) {
			if got := Reverse(tt.args); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
