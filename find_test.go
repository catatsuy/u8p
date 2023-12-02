package u8p_test

import (
	"testing"
	"unicode/utf8"

	"github.com/catatsuy/u8p"
)

func TestFindSuccess(t *testing.T) {
	tests := map[string]struct {
		input  string
		inputL int
		result int
	}{
		"ascii": {
			input:  "hello",
			inputL: 4,
			result: 3,
		},
		"japanese and emoji": {
			input:  "こん😅にちは",
			inputL: 15,
			result: 13,
		},
		"ascii with multi byte glyph 2": {
			input:  "hello🎉",
			inputL: 6,
			result: 5,
		},
		"2 byte string": {
			input:  "ßßßßßßßßßßßßßßß",
			inputL: 13,
			result: 12,
		},
		"empty string": {
			input:  "",
			inputL: 0,
			result: 0,
		},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := u8p.Find(tc.input, tc.inputL)

			if err != nil {
				t.Fatal(err)
			}
			if got != tc.result {
				t.Fatalf("Find(%q) returned %d; expected %d", tc.input, got, tc.result)
			}
		})
	}
}

func TestFindFail(t *testing.T) {
	tests := map[string]struct {
		a       string
		l       int
		want    int
		wantErr bool
		errMsg  string
	}{
		"Empty string":           {a: "", l: 4, want: 0, wantErr: false},
		"Length less than limit": {a: "abc", l: 4, want: 0, wantErr: true, errMsg: "invalid length"},
		"Limit too small":        {a: "abcd", l: 3, want: 0, wantErr: true, errMsg: "l must be greater than 3"},
	}

	for name, tc := range tests {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := u8p.Find(tc.a, tc.l)
			if (err != nil) != tc.wantErr {
				t.Fatalf("Find() error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr && err.Error() != tc.errMsg {
				t.Fatalf("Find() error message = %v, want %v", err.Error(), tc.errMsg)
			}
			if got != tc.want {
				t.Errorf("Find() = %v, want %v", got, tc.want)
			}
		})
	}
}

func FuzzFind(f *testing.F) {
	f.Add(200, "こん😅にちは")
	f.Fuzz(func(t *testing.T, i int, a string) {
		if len(a) == 0 {
			return
		}
		if !utf8.ValidString(a) {
			return
		}
		input := i % len(a)
		if input <= 4 {
			return
		}
		l, err := u8p.Find(a, input)
		if err != nil {
			t.Fatal(err)
		}

		if !utf8.ValidString(a[0:l]) {
			t.Fatalf("invalid utf8")
		}
	})
}
