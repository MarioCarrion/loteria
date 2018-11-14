package loteria

import (
	"math/rand"
	"time"
)

type (
	// Board defines a "tabla", which is 4x4 grid of 16 Cards.
	Board struct {
		marked index
		cards  map[Card]index
	}

	// index indicates the concrete bit to enable in "Board.marked".
	index uint16
)

const (
	// ErrCardNotOnBoard defines the error returned by Mark when the specific
	// Card is not part of the board.
	ErrCardNotOnBoard = Error("card is not on board")
)

// NewBoard returns a new board using concrete cards.
// FIXME validate: all cards on board are built.
func NewBoard(cards []Card) Board {
	board := Board{cards: map[Card]index{}}
	var bit uint16 = 1
	for _, card := range cards {
		board.cards[card] = index(1) << bit
		bit++
	}

	return board
}

// NewRandomBoard returns a board with random Cards.
func NewRandomBoard() Board {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cards := map[Card]index{}
	for len(cards) < 16 {
		v := r.Intn(53)
		if _, ok := cards[Card(v)]; !ok {
			cards[Card(v)] = index(len(cards) + 1)
		}
	}

	return Board{cards: cards}
}

// Mark marks off the card on the board.
func (b *Board) Mark(c Card) error {
	index, ok := b.cards[c]
	if !ok {
		return ErrCardNotOnBoard
	}

	b.marked |= index

	return nil
}
