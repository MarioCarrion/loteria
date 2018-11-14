package loteria

type (
	// Card defines the card, which is part of the board and also announced by
	// the caller.
	Card uint64

	// Deck defines the type containing all the 54 Cards.
	Deck [54]Card
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
)

//nolint
var (
	cardNames = [54]string{
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
	}
)

// NewDeck returns a Deck of sorted Cards.
func NewDeck() Deck {
	r := [54]Card{}
	i := 1
	for i < 54 {
		r[i] = Card(i)
		i++
	}
	return r
}

// String returns the card name.
func (c Card) String() string {
	return cardNames[int(c)]
}
