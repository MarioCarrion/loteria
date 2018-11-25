package loteria

import (
	"fmt"
)

type (

	// Caller defines the actor in charge of announcing the cards.
	Caller struct {
		deck         Deck
		players      map[PlayerName]Board
		boards       map[BoardID]PlayerName
		gameStarted  bool
		gameFinished bool
	}
)

// NewCaller returns a new caller with a shuffled deck of cards.
func NewCaller() Caller {
	deck := NewDeck()
	deck.Shuffle()

	return Caller{
		deck:    deck,
		players: map[PlayerName]Board{},
		boards:  map[BoardID]PlayerName{},
	}
}

// AddPlayer adds the player to the game and it assigns the player a random
// board.
func (c *Caller) AddPlayer(name PlayerName) (Player, error) {
	if c.gameStarted {
		return Player{}, fmt.Errorf("game is in progress, can't add more players")
	}
	if _, ok := c.players[name]; ok {
		return Player{}, fmt.Errorf("player %s is already part of the game", name)
	}

	board := NewRandomBoard()

	if _, ok := c.boards[board.ID()]; ok {
		// XXX retry a few times instead of returning right away
		return Player{}, fmt.Errorf("board was already assigned to another player")
	}

	c.boards[board.ID()] = name
	c.players[name] = board

	return NewPlayer(name, board), nil
}

// Announce announces a card from the deck.
func (c *Caller) Announce() (Card, error) {
	// XXX Must have at least one player to start the game.
	if c.gameFinished {
		return blankCard, fmt.Errorf("game already finished")
	}

	c.gameStarted = true
	card, err := c.deck.Select()
	if err != nil {
		c.gameFinished = true
		return card, err
	}

	// We update our internal boards to use them later in `Loteria` for
	// confirming the player really won.
	for name, board := range c.players {
		if board.Mark(card) == nil {
			c.players[name] = board
		}
	}

	return card, nil
}

// Loteria determines if the player is really the winner
func (c *Caller) Loteria(name PlayerName) error {
	if !c.gameStarted {
		return fmt.Errorf("game has not started yet")
	}
	if c.gameFinished { // XXX consider if name = already winner
		return fmt.Errorf("game already finished")
	}

	board, ok := c.players[name]
	if !ok {
		return fmt.Errorf("player not part of the game")
	}
	if !board.IsWinner() {
		return fmt.Errorf("board is not a winner one")
	}

	c.gameFinished = true
	return nil
}
