package numconv

import (
	"errors"
	"fmt"
	"slices"
	"strings"
)

var (
	cnNumMap = map[rune]int64{
		'零': 0, '一': 1, '二': 2, '三': 3, '四': 4,
		'五': 5, '六': 6, '七': 7, '八': 8, '九': 9,
	}
	cnUnitMap = map[rune]int64{
		'十': 10, '百': 100, '千': 1000,
		'万': 10000, '亿': 100_000_000,
	}
)

var (
	ErrInvalidNumber = errors.New("Number is invalid")
)

// ChineseToArabic convert Chinese numbers to Arabic numbers.
// It only support integers now.
func ChineseToArabic(cn string) (int64, error) {
	if len(cn) < 1 {
		return 0, fmt.Errorf("%w: %s", ErrInvalidNumber, cn)
	}
	// normalize input string
	rn := strings.ReplaceAll(cn, "零", "") // remove zero tokens
	rn = strings.ReplaceAll(rn, "两", "二") // support some words like 两千万

	runes := []rune(rn)
	slices.Reverse(runes)

	var unit int64
	nn := make([]int64, 0, len(runes))
	for _, c := range runes {
		u, ok := cnUnitMap[c]
		if ok {
			if u == 10000 || u == 100000000 {
				nn = append(nn, u)
				unit = 1
			} else {
				unit = u
			}
			continue
		}
		n, ok := cnNumMap[c]
		if !ok {
			return 0, fmt.Errorf("%w: %s", ErrInvalidNumber, cn)
		}
		if unit > 0 {
			n *= unit
			unit = 0
		}
		nn = append(nn, n)
	}

	if unit > 1 {
		nn = append(nn, unit)
	}
	if len(nn) == 1 {
		return nn[0], nil
	}

	var val, tmp int64
	slices.Reverse(nn)
	for _, n := range nn {
		if n == 10000 || n == 100000000 {
			val += tmp * n
			tmp = 0
		} else {
			tmp += n
		}
	}
	val += tmp
	return val, nil
}
