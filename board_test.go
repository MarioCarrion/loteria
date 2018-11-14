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
