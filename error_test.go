package loteria_test

import (
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestError(t *testing.T) {
	tests := [...]struct {
		name            string
		err             error
		expectedMessage string
	}{
		{
			"ErrCardNotOnBoard ",
			loteria.ErrCardNotOnBoard,
			"card is not on board",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(tts *testing.T) {
			if tt.err.Error() != tt.expectedMessage {
				tts.Fatalf("expected %s, actual %s", tt.expectedMessage, tt.err.Error())
			}
		})
	}
}
