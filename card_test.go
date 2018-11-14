package loteria_test

import (
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestNewDeck(t *testing.T) {
	deck := loteria.NewDeck()

	for i, card := range deck {
		t.Run(card.String(), func(tt *testing.T) {
			if uint64(card) > 54 {
				t.Fatalf("invalid Card %d/%d", i, card)
			}
		})
	}
}

func TestCard_String(t *testing.T) {
	// FIXME test ALL values!
	expected := "Harp"
	if loteria.HarpCard.String() != expected {
		t.Fatalf("expected %s, actual %s", expected, loteria.HarpCard.String())
	}
}
