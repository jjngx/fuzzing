package fuzzing_test

import (
	"testing"

	"github.com/jjngx/fuzzing"
)

func TestReverse(t *testing.T) {
	tt := []struct {
		in   string
		want string
	}{
		{"nginx", "xnign"},
		{" ", " "},
		{"123", "321"},
	}

	for _, tc := range tt {
		got := fuzzing.Reverse(tc.in)
		if got != tc.want {
			t.Errorf("got %q, want %q", got, tc.want)
		}
	}
}
