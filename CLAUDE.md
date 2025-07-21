# D&D Turn-Based Card Game - Player State Component

## Project Overview
This is a Go-based dungeons and dragons style turn-based card game focusing on the player state management component. The player state system tracks health, magic, armor, and various temporary/permanent effects applied through game cards.

## Package Structure
```
playerstate/
├── types.go        # Core state types and structs  
├── manager.go      # State manager with main operations
├── effects.go      # Turn-based effects system
├── modifiers.go    # Permanent/semi-permanent modifiers
└── api.go         # Public interfaces and card integration
```

## Key Features
- **Turn-based effect tracking** - Effects expire after specified number of turns
- **Stacking system** - Multiple effects can stack with configurable rules
- **Card integration** - Game cards can apply complex state changes
- **State validation** - Ensures game state remains consistent
- **Change history** - Full audit trail of all state modifications

## Core Concepts

### Effects vs Modifiers
- **Effects**: Temporary changes that expire after N turns (healing spells, poison, shields)
- **Modifiers**: Longer-lasting changes from equipment or permanent spells (armor bonuses, stat increases)

### State Attributes
- Health/MaxHealth - Player hit points
- Mana/MaxMana - Magical energy for spells  
- Armor - Damage reduction
- Turn tracking for effect expiration

## Development Commands

### Testing
```bash
go test ./playerstate/...
```

### Build
```bash
go build ./...
```

### Lint
```bash
golangci-lint run
```

## Architecture Notes
- State changes are atomic and logged for history/replay
- Turn processing handles effect expiration and trigger events
- Card effects can combine multiple state operations in batches
- Public API uses interfaces to allow for testing and future extensibility

## Implementation Status
✅ Package structure designed
✅ Core types defined
✅ Stub functions created with detailed comments
✅ Public API interfaces defined
⏳ Implementation of stub functions (TODO)
⏳ Unit tests (TODO)
⏳ Integration testing (TODO)