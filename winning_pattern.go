package loteria

type (
	// WinningPattern defines the concrete winning pattern.
	WinningPattern uint16
)

//nolint
const (
	FirstRow  WinningPattern = 0xF << 12
	SecondRow WinningPattern = 0xF << 8
	ThirdRow  WinningPattern = 0xF << 4
	FourthRow WinningPattern = 0xF << 0
	//
	FirstColumn  WinningPattern = 0x8888
	SecondColumn WinningPattern = 0x4444
	ThirdColumn  WinningPattern = 0x2222
	FourthColumn WinningPattern = 0x1111
	//
	DiagonalLeftTop  WinningPattern = 0x8421
	DiagonalRightTop WinningPattern = 0x1248
)
