package u8p_test

import (
	"fmt"
	"math/rand/v2"
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

func generateUTF8String(length int) string {
	// ASCII, Cyrillic, Hiragana, and Emoji ranges in Unicode
	var ranges = [][]int{
		{0x0020, 0x007F},   // ASCII
		{0x0400, 0x04FF},   // Cyrillic
		{0x3040, 0x309F},   // Hiragana
		{0x1F600, 0x1F64F}, // Emoticons
	}

	b := make([]rune, length)
	for i := range b {
		r := ranges[rand.IntN(len(ranges))]
		b[i] = rune(r[0] + rand.IntN(r[1]-r[0]+1))
	}
	return string(b)
}

func getByteLengthOfRuneSlice(s string, runeCount int) int {
	runes := []rune(s)
	if runeCount > len(runes) {
		return -1
	}
	return len(string(runes[:runeCount]))
}

func calculateUTF8ByteLengthForRunes(s string, maxRunes int) int {
	var size, totalSize int
	for i := 0; i < maxRunes && len(s) > totalSize; i++ {
		_, size = utf8.DecodeRuneInString(s[totalSize:])
		totalSize += size
	}
	return totalSize
}

func locateRuneAtPosition(s string, position int) int {
	for idx := range s {
		if idx == position {
			return idx
		}
	}

	return -1
}

func BenchmarkFindUTF8Sizes(b *testing.B) {
	sizes := []int{100, 1000, 10000, 100000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			testString := generateUTF8String(size)
			l := len(testString) / 4
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				u8p.Find(testString, l)
			}
		})
	}
}

func BenchmarkGetByteLengthOfRuneSlice(b *testing.B) {
	sizes := []int{100, 500, 1000, 5000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			testString := generateUTF8String(size)
			runeCount := size / 10
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				getByteLengthOfRuneSlice(testString, runeCount)
			}
		})
	}
}

func BenchmarkCalculateUTF8ByteLengthForRunes(b *testing.B) {
	sizes := []int{100, 500, 1000, 5000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			testString := generateUTF8String(size)
			runeCount := size / 10
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				calculateUTF8ByteLengthForRunes(testString, runeCount)
			}
		})
	}
}

func BenchmarkLocateRuneAtPosition(b *testing.B) {
	sizes := []int{100, 500, 1000, 5000, 10000}
	for _, size := range sizes {
		b.Run(fmt.Sprintf("Size%d", size), func(b *testing.B) {
			testString := generateUTF8String(size)
			targetPosition := size / 10
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				locateRuneAtPosition(testString, targetPosition)
			}
		})
	}
}
