package cards

// Card represents a game card with its properties and effects
type Card struct {
	ID          string
	Name        string
	Description string
	Type        CardType
	Cost        int
	Rarity      Rarity
	
	// Effects this card can trigger
	Effects     []CardEffect
	Requirements []Requirement
}

// CardType defines the category of card
type CardType int

const (
	CardTypeSpell CardType = iota
	CardTypeCreature
	CardTypeArtifact
	CardTypeInstant
	CardTypeEnchantment
)

// Rarity defines how rare a card is
type Rarity int

const (
	RarityCommon Rarity = iota
	RarityUncommon
	RarityRare
	RarityLegendary
)

// CardEffect defines what happens when a card is played
type CardEffect struct {
	Type   EffectType
	Target TargetType
	Value  int
	Duration int // For ongoing effects
}

// Deck represents a player's deck of cards
type Deck struct {
	PlayerID string
	Cards    []Card
	Shuffled bool
}

// Hand represents cards currently in a player's hand
type Hand struct {
	PlayerID string
	Cards    []Card
	MaxSize  int
}