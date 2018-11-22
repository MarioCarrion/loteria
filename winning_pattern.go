package loteria

import (
	"fmt"
	"strings"
)

type (
	// WinningPattern defines the concrete winning pattern.
	WinningPattern uint16
)

//nolint
const (
	WinningPatternDefault WinningPattern = 0
	//
	WinningPatternFirstRow  WinningPattern = 0xF << 0
	WinningPatternSecondRow WinningPattern = 0xF << 4
	WinningPatternThirdRow  WinningPattern = 0xF << 8
	WinningPatternFourthRow WinningPattern = 0xF << 12
	//
	WinningPatternFirstColumn  WinningPattern = 0x1111
	WinningPatternSecondColumn WinningPattern = 0x2222
	WinningPatternThirdColumn  WinningPattern = 0x4444
	WinningPatternFourthColumn WinningPattern = 0x8888
	//
	WinningPatternDiagonalRightTop WinningPattern = 0x1248
	WinningPatternDiagonalLeftTop  WinningPattern = 0x8421
)

//nolint
var (
	defaultWinningPatterns = [10]WinningPattern{
		WinningPatternFirstRow, WinningPatternSecondRow, WinningPatternThirdRow, WinningPatternFourthRow,
		WinningPatternFirstColumn, WinningPatternSecondColumn, WinningPatternThirdColumn, WinningPatternFourthColumn,
		WinningPatternDiagonalLeftTop, WinningPatternDiagonalRightTop,
	}
)

// String returns a string value describing the value in a humanized way.
func (w WinningPattern) String() string {
	str := fmt.Sprintf("%016b", w)
	str = strings.Replace(str, "0", ".", -1)
	str = strings.Replace(str, "1", "x", -1)

	// https://github.com/golang/example/blob/master/stringutil/reverse.go
	r := func(s string) string {
		r := []rune(s)
		for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
			r[i], r[j] = r[j], r[i]
		}
		return string(r)
	}

	// Cards on boards are indexed using bit shifting; that's why the rows are
	// in reverse and then the actual columns reversed again.
	return strings.Join([]string{
		r(str[12:16]),
		r(str[8:12]),
		r(str[4:8]),
		r(str[0:4]),
	}, "\n")
}
