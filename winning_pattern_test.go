package loteria_test

import (
	"fmt"
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestWinningPattern_String(t *testing.T) {
	tests := [...]struct {
		name     string
		pattern  loteria.WinningPattern
		expected string
	}{
		{
			"FirstRow",
			loteria.WinningPatternFirstRow,
			"xxxx\n....\n....\n....",
		},
		{
			"SecondRow",
			loteria.WinningPatternSecondRow,
			"....\nxxxx\n....\n....",
		},
		{
			"ThirdRow",
			loteria.WinningPatternThirdRow,
			"....\n....\nxxxx\n....",
		},
		{
			"FourthRow",
			loteria.WinningPatternFourthRow,
			"....\n....\n....\nxxxx",
		},
		{
			"FirstColumn",
			loteria.WinningPatternFirstColumn,
			"x...\nx...\nx...\nx...",
		},
		{
			"SecondColumn",
			loteria.WinningPatternSecondColumn,
			".x..\n.x..\n.x..\n.x..",
		},
		{
			"ThirdColumn",
			loteria.WinningPatternThirdColumn,
			"..x.\n..x.\n..x.\n..x.",
		},
		{
			"FourthColumn",
			loteria.WinningPatternFourthColumn,
			"...x\n...x\n...x\n...x",
		},
		{
			"DiagonalLeftTop",
			loteria.WinningPatternDiagonalLeftTop,
			"x...\n.x..\n..x.\n...x",
		},
		{
			"DiagonalRightTop",
			loteria.WinningPatternDiagonalRightTop,
			"...x\n..x.\n.x..\nx...",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(tts *testing.T) {
			if got := tt.pattern.String(); got != tt.expected {
				t.Fatalf("expected:\n%s\ngot:\n%s", tt.expected, got)
			}
		})
		fmt.Println(tt.pattern.String())
	}
}
