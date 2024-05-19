package u8p

import (
	"fmt"
	"unicode/utf8"
)

func Find(a string, l int) (int, error) {
	if len(a) == 0 {
		return 0, nil
	}
	if l <= 3 {
		return 0, fmt.Errorf("l must be greater than 3")
	}
	if len(a) <= l {
		return 0, fmt.Errorf("invalid length")
	}
	for i := l - 1; i >= l-4; i-- {
		if utf8.RuneStart(a[i]) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("invalid utf8")
}
