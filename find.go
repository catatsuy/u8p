package u8p

import "fmt"

const (
	t1 = 0b00000000
	tx = 0b10000000
	t2 = 0b00000110
	t3 = 0b00001110
	t4 = 0b00011110
)

func Find(a string, l int) (int, error) {
	if len(a) == 0 || len(a) <= l {
		return 0, fmt.Errorf("invalid length")
	}
	if l <= 4 {
		return 0, fmt.Errorf("invalid length")
	}
	for i := l - 1; i >= l-4; i-- {
		if isUTF8LeadByte(a[i]) {
			return i, nil
		}
	}
	return 0, fmt.Errorf("invalid utf8")
}

func isUTF8LeadByte(tmp byte) bool {
	if tmp&tx == t1 {
		return true
	}
	tmp >>= 3
	if tmp == t4 {
		return true
	}
	tmp >>= 1
	if tmp == t3 {
		return true
	}
	tmp >>= 1
	if tmp == t2 {
		return true
	}
	return false
}
