package loteria

import (
	"crypto/md5" //nolint: gosec
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type (
	// Board defines a "tabla", which is 4x4 grid of 16 Cards.
	Board struct {
		WinningPattern WinningPattern
		marked         index
		cards          map[Card]index
		id             BoardID
	}

	// BoardID represents the Board ID
	BoardID string

	// index indicates the concrete bit to enable in "Board.marked".
	index uint16
)

const (
	// ErrCardNotOnBoard defines the error returned by Mark when the specific
	// Card is not part of the board.
	ErrCardNotOnBoard = Error("card is not on board")
)

// NewBoard returns a new board using concrete cards.
// FIXME validate: cards uniqueness.
func NewBoard(cards []Card) Board {
	board := Board{cards: map[Card]index{}}
	var bit uint16
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
			cards[Card(v)] = index(1) << uint16(len(cards))
		}
	}

	return Board{cards: cards}
}

// Cards returns the cards on the board.
func (b *Board) Cards() [16]Card {
	findIndex := func(n uint16) int {
		var i uint16 = 1
		var pos uint16 = 1

		for i&n == 0 {
			i = i << 1
			pos++
		}

		return int(pos - 1)
	}

	cards := [16]Card{}

	for k, v := range b.cards {
		cards[findIndex(uint16(v))] = k
	}

	return cards
}

// ID returns the Board Identifier
func (b *Board) ID() BoardID {
	if b.id != "" {
		return b.id
	}

	a := []byte{}
	for _, c := range b.cards {
		jb, _ := json.Marshal(c) //nolint: gosec
		a = append(a, jb...)
	}
	b.id = BoardID(fmt.Sprintf("%x", md5.Sum(a))) //nolint: gosec

	return b.id
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
