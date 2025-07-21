package game

import "time"

// Game represents the overall game state and manages turns
type Game struct {
	ID          string
	Players     []string
	CurrentTurn int
	Phase       GamePhase
	Winner      string
	Status      GameStatus
	CreatedAt   time.Time
	
	// Turn management
	ActivePlayer string
	TurnOrder    []string
	TurnLimit    int
}

// GamePhase represents the current phase of a turn
type GamePhase int

const (
	PhaseStart GamePhase = iota
	PhaseMain
	PhaseCombat
	PhaseEnd
)

// GameStatus represents the current state of the game
type GameStatus int

const (
	StatusWaiting GameStatus = iota
	StatusActive
	StatusPaused
	StatusFinished
	StatusAbandoned
)

// TurnAction represents an action taken during a turn
type TurnAction struct {
	PlayerID  string
	ActionType ActionType
	CardID    string
	TargetID  string
	Timestamp time.Time
}

// ActionType defines types of actions players can take
type ActionType int

const (
	ActionPlayCard ActionType = iota
	ActionAttack
	ActionEndTurn
	ActionUseAbility
	ActionEquipItem
)