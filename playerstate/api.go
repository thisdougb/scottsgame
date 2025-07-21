package playerstate

// StateManager defines the public interface for player state management
type StateManager interface {
	// Player management
	CreatePlayer(playerID string, maxHealth, maxMana int) *PlayerState
	GetPlayer(playerID string) (*PlayerState, error)
	
	// Core state operations
	ApplyDamage(playerID string, damage int, source string) error
	ApplyHealing(playerID string, healing int, source string) error
	ConsumeResource(playerID string, resourceType string, amount int) error
	
	// Turn management
	ProcessTurnStart(playerID string, turn int) error
	ProcessTurnEnd(playerID string, turn int) error
	
	// Effect management
	AddEffect(playerID string, effect Effect) error
	RemoveEffect(playerID string, effectID string) error
	RemoveEffectsBySource(playerID string, source string) error
	GetActiveEffects(playerID string) ([]Effect, error)
	
	// Modifier management
	AddModifier(playerID string, modifier Modifier) error
	RemoveModifier(playerID string, modifierID string) error
	RemoveModifiersBySource(playerID string, source string) error
	GetActiveModifiers(playerID string) ([]Modifier, error)
	
	// State querying
	GetEffectiveValue(playerID string, attribute string) (int, error)
	GetStateHistory(playerID string) ([]StateChange, error)
	ValidateState(playerID string) error
}

// Ensure Manager implements StateManager interface
var _ StateManager = (*Manager)(nil)

// CardEffect represents the state changes that a game card can apply
type CardEffect struct {
	CardID      string
	Name        string
	Description string
	
	// Direct state changes
	HealthChange int
	ManaChange   int
	ArmorChange  int
	
	// Effects to apply
	EffectsToAdd []Effect
	
	// Modifiers to apply  
	ModifiersToAdd []Modifier
	
	// Targeting
	TargetSelf   bool
	TargetEnemy  bool
	RequiresTarget bool
}

// ApplyCardEffect applies all effects from a card to the target player(s)
func (m *Manager) ApplyCardEffect(cardEffect CardEffect, casterID, targetID string) error {
	// TODO: Apply all effects defined in the card effect
	// Handle direct stat changes (damage, healing, mana changes)
	// Apply any effects and modifiers specified by the card
	// Validate targeting rules (self, enemy, requires target)
	// Log all state changes with card as source
	panic("not implemented")
}

// GetPlayerSummary returns a summary of player state suitable for UI display
func (m *Manager) GetPlayerSummary(playerID string) (*PlayerSummary, error) {
	// TODO: Generate a summary with current and effective values
	// Include active effects count, status condition descriptions
	// Format for easy display in game UI
	panic("not implemented")
}

// PlayerSummary contains formatted player state information for UI
type PlayerSummary struct {
	PlayerID string
	
	// Current values
	Health    int
	MaxHealth int
	Mana      int
	MaxMana   int
	Armor     int
	
	// Status indicators
	StatusEffects []string // Human-readable effect descriptions
	ActiveEffectsCount int
	ActiveModifiersCount int
	
	// Turn info
	CurrentTurn int
	
	// Warnings/alerts
	LowHealth bool
	LowMana   bool
	HasNegativeEffects bool
}

// BatchOperation allows multiple state changes to be applied atomically
type BatchOperation struct {
	PlayerID string
	Operations []StateOperation
}

// StateOperation represents a single state change operation
type StateOperation struct {
	Type   OperationType
	Params map[string]interface{}
}

// OperationType defines types of state operations that can be batched
type OperationType int

const (
	OpApplyDamage OperationType = iota
	OpApplyHealing
	OpAddEffect
	OpAddModifier
	OpConsumeResource
)

// ExecuteBatch applies multiple state operations atomically
func (m *Manager) ExecuteBatch(batch BatchOperation) error {
	// TODO: Apply all operations in the batch atomically
	// If any operation fails, roll back all changes
	// Useful for complex card effects that do multiple things
	// Log all changes as a single transaction
	panic("not implemented")
}