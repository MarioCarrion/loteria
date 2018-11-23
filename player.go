package loteria

type (
	// PlayerName defines the name of the player, it's used to identify the user
	// uniquely.
	PlayerName string

	// Player defines the person in the game with a Board ready to play.
	Player struct {
		board Board
		name  PlayerName
	}
)

// NewPlayer returns a new Player.
func NewPlayer(name PlayerName, board Board) Player {
	return Player{
		name:  name,
		board: board,
	}
}

// IsWinner indicates whether the marked cards win the game.
func (p *Player) IsWinner() bool {
	return p.board.IsWinner()
}

// Mark marks off the card on the board.
func (p *Player) Mark(c Card) error {
	return p.board.Mark(c)
}

// Name returns the player name.
func (p *Player) Name() PlayerName {
	return p.name
}
