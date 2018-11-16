package loteria

type (
	// WinningPattern defines the concrete winning pattern.
	WinningPattern uint16
)

//nolint
const (
	NoWinner WinningPattern = 0
	//
	FirstRow  WinningPattern = 0xF << 0
	SecondRow WinningPattern = 0xF << 4
	ThirdRow  WinningPattern = 0xF << 8
	FourthRow WinningPattern = 0xF << 12
	//
	FirstColumn  WinningPattern = 0x1111
	SecondColumn WinningPattern = 0x2222
	ThirdColumn  WinningPattern = 0x4444
	FourthColumn WinningPattern = 0x8888
	//
	DiagonalRightTop WinningPattern = 0x1248
	DiagonalLeftTop  WinningPattern = 0x8421
)

//nolint
var (
	defaultWinningPatterns = [10]WinningPattern{
		FirstRow, SecondRow, ThirdRow, FourthRow,
		FirstColumn, SecondColumn, ThirdColumn, FourthColumn,
		DiagonalLeftTop, DiagonalRightTop,
	}
)
