package loteria

import (
	"fmt"
	"math/rand"
	"time"
)

type (
	// Card defines the card, which is part of the Boards and also announced by
	// the caller.
	Card uint64

	// Deck defines the type containing all the 54 Cards.
	Deck struct {
		cards       [deckLength]Card
		selectIndex int
	}
)

const (
	deckLength = 54
)

//nolint
const (
	RoosterCard Card = iota
	DevilCard
	LadyCard
	DandyCard
	UmbrellaCard
	MermaidCard
	LadderCard
	BottleCard
	BarrelCard
	TreeCard
	MelonCard
	BraveManCard
	BonnetCard
	DeathCard
	PearCard
	FlagCard
	MandolinCard
	CelloCard
	HeronCard
	BirdCard
	HandCard
	BootCard
	MoonCard
	ParrotCard
	DrunkardCard
	BussinessManCard
	HeartCard
	WatermelonCard
	DrumCard
	ShrimpCard
	ArrowsCard
	MusicianCard
	SpiderCard
	SoldierCard
	StarCard
	SaucepanCard
	WorldCard
	FeatherCard
	CactusCard
	ScorpionCard
	RoseCard
	SkullCard
	BellCard
	WaterPitcherCard
	DeerCard
	SunCard
	CrownCard
	CanoeCard
	PineTreeCard
	FishCard
	PalmTreeCard
	FlowerPotCard
	HarpCard
	FrogCard
	//
	blankCard
)

//nolint
var (
	cardNames = [(deckLength + 1)]string{
		"Rooster",
		"Devil",
		"Lady",
		"Dandy",
		"Umbrella",
		"Mermaid",
		"Ladder",
		"Bottle",
		"Barrel",
		"Tree",
		"Melon",
		"BraveMan",
		"Bonnet",
		"Death",
		"Pear",
		"Flag",
		"Mandolin",
		"Cello",
		"Heron",
		"Bird",
		"Hand",
		"Boot",
		"Moon",
		"Parrot",
		"Drunkard",
		"BussinessMan",
		"Heart",
		"Watermelon",
		"Drum",
		"Shrimp",
		"Arrows",
		"Musician",
		"Spider",
		"Soldier",
		"Star",
		"Saucepan",
		"World",
		"Feather",
		"Cactus",
		"Scorpion",
		"Rose",
		"Skull",
		"Bell",
		"WaterPitcher",
		"Deer",
		"Sun",
		"Crown",
		"Canoe",
		"PineTree",
		"Fish",
		"PalmTree",
		"FlowerPot",
		"Harp",
		"Frog",
		//
		"Blank",
	}
)

// NewDeck returns a Deck with sorted Cards.
func NewDeck() Deck {
	r := [deckLength]Card{}

	for i := 0; i < deckLength; i++ {
		r[i] = Card(i)
	}

	return Deck{cards: r}
}

// String returns the card name.
func (c Card) String() string {
	return cardNames[int(c)]
}

// Select select a card from the deck
func (d *Deck) Select() (Card, error) {
	if d.selectIndex == deckLength {
		return blankCard, fmt.Errorf("deck is empty")
	}

	r := d.cards[d.selectIndex]
	d.selectIndex++
	return r, nil
}

// Shuffle shuffles the deck.
func (d *Deck) Shuffle() {
	d.selectIndex = 0
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
}
