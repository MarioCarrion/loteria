package loteria_test

import (
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestBoard_Mark(t *testing.T) {
	tests := [...]struct {
		name          string
		board         loteria.Board
		input         loteria.Card
		expectedError error
	}{
		{
			"OK",
			loteria.NewBoard([]loteria.Card{loteria.FrogCard, loteria.DevilCard}),
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
	cards := []loteria.Card{
		loteria.DeathCard, loteria.FlagCard, loteria.MoonCard, loteria.DrumCard,
		loteria.SpiderCard, loteria.SkullCard, loteria.FrogCard, loteria.LadderCard,
		loteria.BonnetCard, loteria.BirdCard, loteria.SoldierCard, loteria.MermaidCard,
		loteria.RoosterCard, loteria.DrunkardCard, loteria.FeatherCard, loteria.CactusCard,
	}

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
			loteria.FirstRow,
		},
		{
			"SecondRow",
			[4]loteria.Card{loteria.SpiderCard, loteria.SkullCard, loteria.FrogCard, loteria.LadderCard},
			true,
			loteria.SecondRow,
		},
		{
			"ThirdRow",
			[4]loteria.Card{loteria.BonnetCard, loteria.BirdCard, loteria.SoldierCard, loteria.MermaidCard},
			true,
			loteria.ThirdRow,
		},
		{
			"FourthRow",
			[4]loteria.Card{loteria.RoosterCard, loteria.DrunkardCard, loteria.FeatherCard, loteria.CactusCard},
			true,
			loteria.FourthRow,
		},
		{
			"FirstColumn",
			[4]loteria.Card{loteria.DeathCard, loteria.SpiderCard, loteria.BonnetCard, loteria.RoosterCard},
			true,
			loteria.FirstColumn,
		},
		{
			"SecondColumn",
			[4]loteria.Card{loteria.FlagCard, loteria.SkullCard, loteria.BirdCard, loteria.DrunkardCard},
			true,
			loteria.SecondColumn,
		},
		{
			"ThirdColumn",
			[4]loteria.Card{loteria.MoonCard, loteria.FrogCard, loteria.SoldierCard, loteria.FeatherCard},
			true,
			loteria.ThirdColumn,
		},
		{
			"DiagonalLeftTop",
			[4]loteria.Card{loteria.DeathCard, loteria.SkullCard, loteria.SoldierCard, loteria.CactusCard},
			true,
			loteria.DiagonalLeftTop,
		},
		{
			"DiagonalRightTop ",
			[4]loteria.Card{loteria.DrumCard, loteria.FrogCard, loteria.BirdCard, loteria.RoosterCard},
			true,
			loteria.DiagonalRightTop,
		},
		{
			"NoWinner",
			[4]loteria.Card{loteria.DrumCard},
			false,
			loteria.NoWinner,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			board := loteria.NewBoard(cards)
			for _, card := range tt.cards {
				board.Mark(card)
			}

			if winner := board.Winner(); winner != tt.expectedWinner {
				t.Fatalf("expected winner to be %T, actual %T", tt.expectedWinner, winner)
			}
			if board.WinningPattern != tt.expectedWinningPattern {
				t.Fatalf("expected winner to be %016b, actual %016b", tt.expectedWinningPattern, board.WinningPattern)
			}
		})
	}
}
