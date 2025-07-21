package playerstate

import "time"

// PlayerState represents the current state of a player character
type PlayerState struct {
	PlayerID string
	
	// Core attributes
	Health    int
	MaxHealth int
	Mana      int
	MaxMana   int
	Armor     int
	
	// Active effects and modifiers
	ActiveEffects []Effect
	Modifiers     []Modifier
	
	// Turn tracking
	CurrentTurn int
	LastUpdated time.Time
}

// Effect represents a temporary state change that expires after a certain number of turns
type Effect struct {
	ID          string
	Name        string
	Description string
	Type        EffectType
	Value       int
	Duration    int        // turns remaining
	AppliedTurn int        // turn when applied
	Source      string     // card or source that applied this effect
	Stackable   bool       // whether multiple instances can stack
}

// Modifier represents a persistent state change from equipment, spells, etc.
type Modifier struct {
	ID          string
	Name        string
	Type        ModifierType
	Value       int
	Source      string
	Permanent   bool
	AppliedTurn int
}

// EffectType defines the type of effect applied to player state
type EffectType int

const (
	EffectTypeHealthBonus EffectType = iota
	EffectTypeManaBonus
	EffectTypeArmorBonus
	EffectTypeDamageReduction
	EffectTypeHealOverTime
	EffectTypeDamageOverTime
	EffectTypeShield
	EffectTypeRegeneration
)

// ModifierType defines the type of modifier applied to player state
type ModifierType int

const (
	ModifierTypeHealth ModifierType = iota
	ModifierTypeMana
	ModifierTypeArmor
	ModifierTypeDamage
	ModifierTypeResistance
)

// StateChange represents a single change to player state
type StateChange struct {
	Type        ChangeType
	Attribute   string
	OldValue    int
	NewValue    int
	Source      string
	Turn        int
	Timestamp   time.Time
}

// ChangeType defines the type of state change
type ChangeType int

const (
	ChangeTypeDirect ChangeType = iota
	ChangeTypeEffect
	ChangeTypeModifier
	ChangeTypeExpiration
)