package fuzzing_test

import (
	"testing"
	"unicode/utf8"

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

func FuzzReverse(f *testing.F) {
	tt := []string{"nginx", " ", "123", "XYZ", "!alf@"}
	for _, tc := range tt {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, input string) {
		rev := fuzzing.Reverse(input)
		doubleRev := fuzzing.Reverse(rev)
		t.Logf("Number of runes: orig=%d, rev=%d, doubleRev=%d",
			utf8.RuneCountInString(input),
			utf8.RuneCountInString(rev),
			utf8.RuneCountInString(doubleRev),
		)

		if input != doubleRev {
			t.Errorf("Before: %q, after: %q", input, doubleRev)
		}
		if utf8.ValidString(input) && !utf8.ValidString(rev) {
			t.Errorf("got invalid utf8 string: %q", rev)
		}
	})
}
