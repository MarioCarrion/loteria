package loteria_test

import (
	"reflect"
	"testing"

	"github.com/MarioCarrion/loteria"
)

func TestDeck_Select(t *testing.T) {
	deck := loteria.NewDeck()

	i := 0
	for {
		_, err := deck.Select()
		i++
		if err != nil {
			break
		}
	}
	if i != 55 {
		t.Fatalf("expected to return error when deck is empty")
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
