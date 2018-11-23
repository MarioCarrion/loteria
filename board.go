package loteria

import (
	"math/rand"
	"time"
)

type (
	// Board defines a "tabla", which is 4x4 grid of 16 Cards.
	Board struct {
		WinningPattern WinningPattern
		id             BoardID
		marked         boardIndex
		cardsComputed  bool
		cardsMap       map[Card]boardIndex
		cards          [16]Card
	}

	// BoardID represents the Board ID
	BoardID uint16

	// boardIndex indicates the location (bitwise) of cards on the board.
	boardIndex uint16
)

const (
	// ErrCardNotOnBoard defines the error returned by Mark when the specific
	// Card is not part of the board.
	ErrCardNotOnBoard = Error("card is not on board")
)

// NewBoard returns a new board using concrete cards.
// FIXME validate: cards uniqueness.
func NewBoard(cards [16]Card) Board {
	board := Board{cardsMap: map[Card]boardIndex{}}
	var bit uint16
	for _, card := range cards {
		board.cardsMap[card] = boardIndex(1) << bit
		bit++
	}

	return board
}

// NewRandomBoard returns a board with random Cards.
func NewRandomBoard() Board {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	cards := map[Card]boardIndex{}
	for len(cards) < 16 {
		v := r.Intn(53)
		if _, ok := cards[Card(v)]; !ok {
			cards[Card(v)] = boardIndex(1) << uint16(len(cards))
		}
	}

	return Board{cardsMap: cards}
}

// Cards returns the cards on the board.
func (b *Board) Cards() [16]Card {
	if b.cardsComputed {
		return b.cards
	}
	b.cardsComputed = true

	findIndex := func(n uint16) int {
		var i uint16 = 1
		var pos uint16 = 1

		for i&n == 0 {
			i = i << 1
			pos++
		}

		return int(pos - 1)
	}

	for k, v := range b.cardsMap {
		b.cards[findIndex(uint16(v))] = k
	}

	return b.cards
}

// ID returns the Board Identifier
func (b *Board) ID() BoardID {
	if b.id != 0 {
		return b.id
	}

	var res uint16
	for _, c := range b.cardsMap {
		res |= uint16(c)
	}

	b.id = BoardID(res)

	return b.id
}

// Mark marks off the card on the board.
func (b *Board) Mark(c Card) error {
	index, ok := b.cardsMap[c]
	if !ok {
		return ErrCardNotOnBoard
	}

	b.marked |= index

	return nil
}

// IsWinner indicates whether the marked cards win the game.
func (b *Board) IsWinner() bool {
	for _, pattern := range defaultWinningPatterns {
		if (uint16(b.marked) & uint16(pattern)) == uint16(pattern) {
			b.WinningPattern = pattern
			return true
		}
	}
	return false
}
