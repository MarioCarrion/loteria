package main

import (
	"fmt"
)

type (
	Board struct {
		selected uint16
	}
)

const (
	FirstRow  uint16 = 0xF << 12
	SecondRow uint16 = 0xF << 8
	ThirdRow  uint16 = 0xF << 4
	FourthRow uint16 = 0xF << 0
	//
	FirstColumn  uint16 = 0x8888
	SecondColumn uint16 = 0x4444
	ThirdColumn  uint16 = 0x2222
	FourthColumn uint16 = 0x1111
	//
	DiagonalOne uint16 = 0x8421
	DiagonalTwo uint16 = 0x1248
)

func main() {
	//  1, 2, 3, 4
	//  5, 6, 7, 8
	//  9,10,11,12
	// 13,14,15,16

	winners := [...]uint16{
		FirstRow, SecondRow, ThirdRow, FourthRow,
		FirstColumn, SecondColumn, ThirdColumn, FourthColumn,
		DiagonalOne, DiagonalTwo,
	}

	for _, w := range winners {
		str := fmt.Sprintf("%016b\n", w)
		fmt.Printf("%s\n", str[0:4])
		fmt.Printf("%s\n", str[4:8])
		fmt.Printf("%s\n", str[8:12])
		fmt.Printf("%s\n-----\n", str[12:16])
	}
}
