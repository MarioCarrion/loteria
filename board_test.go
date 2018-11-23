package loteria_test

import (
	"reflect"
	"testing"

	"github.com/MarioCarrion/loteria"
)

func newTestCards() [16]loteria.Card {
	return [16]loteria.Card{
		loteria.DeathCard, loteria.FlagCard, loteria.MoonCard, loteria.DrumCard,
		loteria.SpiderCard, loteria.SkullCard, loteria.FrogCard, loteria.LadderCard,
		loteria.BonnetCard, loteria.BirdCard, loteria.SoldierCard, loteria.MermaidCard,
		loteria.RoosterCard, loteria.DrunkardCard, loteria.FeatherCard, loteria.CactusCard,
	}
}

func TestBoard_Cards(t *testing.T) {
	cards := newTestCards()

	board := loteria.NewBoard(cards)
	got := board.Cards()
	if !reflect.DeepEqual(cards, got) {
		t.Fatalf("expected:\n%s\ngot:\n%s", cards, got)
	}

	got1 := board.Cards() // memoized
	if !reflect.DeepEqual(cards, got1) {
		t.Fatalf("expected:\n%s\ngot:\n%s", cards, got1)
	}
}

func TestBoard_ID(t *testing.T) {
	board := loteria.NewRandomBoard()

	expected := board.ID()
	got := board.ID()
	if expected != got {
		t.Fatalf("expected %d, got %d", expected, got)
	}

	cards := board.Cards()
	newBoard := loteria.NewBoard(cards)
	got = newBoard.ID()
	if expected != got {
		t.Fatalf("expected %d, got %d", expected, got)
	}
}

func TestBoard_Mark(t *testing.T) {
	tests := [...]struct {
		name          string
		board         loteria.Board
		input         loteria.Card
		expectedError error
	}{
		{
			"OK",
			loteria.NewBoard([16]loteria.Card{loteria.FrogCard, loteria.DevilCard}),
			loteria.FrogCard,
			nil,
		},
		{
			"ErrCardNotOnBoard",
			loteria.Board{},
			loteria.FrogCard,
			loteria.ErrCardNotOnBoard,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.board.Mark(tt.input)
			if e != tt.expectedError {
				t.Fatalf("expected %s, got %s", tt.expectedError, e)
			}
		})
	}
}

func TestBoard_Winner(t *testing.T) {
	cards := newTestCards()

	tests := [...]struct {
		name                   string
		cards                  [4]loteria.Card
		expectedWinner         bool
		expectedWinningPattern loteria.WinningPattern
	}{
		{
			"FirstRow",
			[4]loteria.Card{loteria.DeathCard, loteria.FlagCard, loteria.MoonCard, loteria.DrumCard},
			true,
			loteria.WinningPatternFirstRow,
		},
		{
			"SecondRow",
			[4]loteria.Card{loteria.SpiderCard, loteria.SkullCard, loteria.FrogCard, loteria.LadderCard},
			true,
			loteria.WinningPatternSecondRow,
		},
		{
			"ThirdRow",
			[4]loteria.Card{loteria.BonnetCard, loteria.BirdCard, loteria.SoldierCard, loteria.MermaidCard},
			true,
			loteria.WinningPatternThirdRow,
		},
		{
			"FourthRow",
			[4]loteria.Card{loteria.RoosterCard, loteria.DrunkardCard, loteria.FeatherCard, loteria.CactusCard},
			true,
			loteria.WinningPatternFourthRow,
		},
		{
			"FirstColumn",
			[4]loteria.Card{loteria.DeathCard, loteria.SpiderCard, loteria.BonnetCard, loteria.RoosterCard},
			true,
			loteria.WinningPatternFirstColumn,
		},
		{
			"SecondColumn",
			[4]loteria.Card{loteria.FlagCard, loteria.SkullCard, loteria.BirdCard, loteria.DrunkardCard},
			true,
			loteria.WinningPatternSecondColumn,
		},
		{
			"ThirdColumn",
			[4]loteria.Card{loteria.MoonCard, loteria.FrogCard, loteria.SoldierCard, loteria.FeatherCard},
			true,
			loteria.WinningPatternThirdColumn,
		},
		{
			"DiagonalLeftTop",
			[4]loteria.Card{loteria.DeathCard, loteria.SkullCard, loteria.SoldierCard, loteria.CactusCard},
			true,
			loteria.WinningPatternDiagonalLeftTop,
		},
		{
			"DiagonalRightTop",
			[4]loteria.Card{loteria.DrumCard, loteria.FrogCard, loteria.BirdCard, loteria.RoosterCard},
			true,
			loteria.WinningPatternDiagonalRightTop,
		},
		{
			"DefaultWinningPattern",
			[4]loteria.Card{loteria.DrumCard},
			false,
			loteria.WinningPatternDefault,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := loteria.NewBoard(cards)
			for _, card := range tt.cards {
				board.Mark(card)
			}

			if winner := board.IsWinner(); winner != tt.expectedWinner {
				t.Fatalf("expected winner to be %T, actual %T", tt.expectedWinner, winner)
			}
			if board.WinningPattern != tt.expectedWinningPattern {
				t.Fatalf("expected winner to be %016b, actual %016b", tt.expectedWinningPattern, board.WinningPattern)
			}
		})
	}
}
