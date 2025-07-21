package events

import "time"

// EventSystem manages game events and notifications
type EventSystem struct {
	subscribers map[EventType][]EventHandler
	history     []Event
}

// Event represents something that happened in the game
type Event struct {
	ID        string
	Type      EventType
	Source    string
	Target    string
	Data      map[string]interface{}
	Timestamp time.Time
}

// EventType defines categories of game events
type EventType int

const (
	EventPlayerCreated EventType = iota
	EventCardPlayed
	EventDamageDealt
	EventHealthChanged
	EventEffectApplied
	EventEffectExpired
	EventTurnStarted
	EventTurnEnded
	EventGameWon
	EventGameLost
	EventItemEquipped
	EventItemUnequipped
)

// EventHandler is a function that processes events
type EventHandler func(event Event) error

// GameLog represents a complete record of game events
type GameLog struct {
	GameID string
	Events []Event
	StartTime time.Time
	EndTime   time.Time
}