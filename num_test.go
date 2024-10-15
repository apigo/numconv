package numconv

import "testing"

func TestChineseToArabic(t *testing.T) {
	tests := []struct {
		Input   string
		Expect  int64
		Invalid bool
	}{
		{"", 0, true},
		{"数", 0, true},
		{"三百个", 0, true},
		{"一", 1, false},
		{"二", 2, false},
		{"十", 10, false},
		{"百", 100, false},
		{"千", 1000, false},
		{"万", 10000, false},
		{"亿", 100000000, false},
		{"一万", 10000, false},
		{"三万", 30000, false},
		{"一十", 10, false},
		{"一十三", 13, false},
		{"十二", 12, false},
		{"六十二", 62, false},
		{"六百四十五", 645, false},
		{"两百", 200, false},
		{"四百零八", 408, false},
		{"三千七百二十一", 3721, false},
		{"一亿三千七百二十一", 100003721, false},
		{"九十二万零七十五", 920075, false},
	}

	for _, tc := range tests {
		n, err := ChineseToArabic(tc.Input)
		if err != nil {
			if tc.Invalid {
				continue
			}
			t.Fatalf("Unexpect error for parse valid number: %s", tc.Input)
		}
		if n != tc.Expect {
			t.Fatalf("Convert %s failed, expect: %d, got: %d", tc.Input, tc.Expect, n)
		}
	}
}
