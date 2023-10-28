package u8p_test

import (
	"testing"
	"unicode/utf8"

	. "github.com/catatsuy/u8p"
)

func TestFind(t *testing.T) {
	a := "こん😅にちは"
	l, err := Find(a, 15)
	if err != nil {
		t.Fatal(err)
	}

	expected := 13

	if l != expected {
		t.Fatalf("want %d, got %d", expected, l)
	}
}

func FuzzFind(f *testing.F) {
	f.Add(200, "こん😅にちは")
	f.Fuzz(func(t *testing.T, i int, a string) {
		if len(a) <= 2 {
			return
		}
		if i < 100 || len(a) <= 100 {
			return
		}
		if i%len(a) <= 4 {
			return
		}
		if !utf8.ValidString(a) {
			return
		}
		l, err := Find(a, i%len(a))
		if err != nil {
			t.Fatal(err)
		}

		if !utf8.ValidString(a[0:l]) {
			t.Fatalf("invalid utf8")
		}
	})
}
