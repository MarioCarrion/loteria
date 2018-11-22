package loteria_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestDeck_Select(t *testing.T) {
	deck := loteria.NewDeck()

	i := 0
	for {
		t.Run(fmt.Sprintf("Index %d", i), func(tt *testing.T) {
			card, err := deck.Select()
			if err != nil {
				if i != 53 {
					tt.Fatalf("expected to return error when deck is empty")
				}
				tt.Fatalf("unexpected error: %s", err)
			}

			if uint64(card) >= 54 {
				tt.Fatalf("invalid Card %d/%d", i, card)
			}
		})
		i++
		if i >= 54 {
			break
		}
	}
}

func TestCard_String(t *testing.T) {
	// FIXME test ALL values!
	expected := "Harp"
	if loteria.HarpCard.String() != expected {
		t.Fatalf("expected %s, actual %s", expected, loteria.HarpCard.String())
	}
}

func TestDeck_Shuffle(t *testing.T) {
	orig := loteria.NewDeck()
	changed := orig

	if !reflect.DeepEqual(orig, changed) {
		t.Fatalf("expected to be equal")
	}

	orig.Shuffle()

	if reflect.DeepEqual(orig, changed) {
		t.Fatalf("expected to be different")
	}
}
