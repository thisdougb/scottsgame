# D&D Turn-Based Card Game

## Project Overview
This is a Go-based dungeons and dragons style turn-based card game with a modular architecture. The game features player state management, card systems, combat mechanics, inventory management, and AI opponents, all connected through an event-driven architecture.

## Project Structure
```
├── ARCHITECTURE.md         # Overall system architecture
├── CLAUDE.md              # This file - project documentation
├── playerstate/           # Player state management
│   ├── types.go          # Core state types and structs  
│   ├── manager.go        # State manager with main operations
│   ├── effects.go        # Turn-based effects system
│   ├── modifiers.go      # Permanent/semi-permanent modifiers
│   ├── api.go           # Public interfaces and card integration
│   └── README.md         # Player state component documentation
├── game/                  # Game state and turn management
│   └── types.go          # Game, turn, and action types
├── cards/                 # Card system and deck management
│   └── types.go          # Card definitions and effects
├── combat/                # Battle mechanics and damage calculation
│   └── types.go          # Combat resolution types
├── inventory/             # Items and equipment management
│   └── types.go          # Item and equipment types
├── events/                # Event system for component communication
│   └── types.go          # Event types and logging
└── ai/                    # Computer player AI
    └── types.go          # AI decision making types
```

## Core Game Components

### Player State System (`playerstate/`)
- Tracks health, mana, armor, and other character attributes
- Manages temporary effects that expire after N turns
- Handles permanent modifiers from equipment and spells
- Provides full audit trail of all state changes
- **Documentation**: See `playerstate/README.md` for detailed design

### Game Engine (`game/`)
- Manages overall game state and turn order
- Coordinates between all game components
- Validates player actions and game rules
- Tracks game phases and win/loss conditions

### Card System (`cards/`)
- Defines card types, effects, and metadata
- Manages player decks and hands
- Integrates with player state for effect application
- Supports various card types (spells, creatures, artifacts)

### Combat System (`combat/`)
- Resolves battles between players and creatures
- Calculates damage with armor and resistance
- Manages combat phases and targeting
- Integrates with player state for damage application

### Inventory System (`inventory/`)
- Manages player items and equipment
- Tracks equipment bonuses and effects
- Handles inventory capacity and item stacking
- Integrates with player state for stat modifications

### Event System (`events/`)
- Provides pub/sub messaging between components
- Logs all game actions for replay and debugging
- Enables loose coupling between game systems
- Supports spectator mode and game analysis

### AI System (`ai/`)
- Implements computer opponents with various strategies
- Makes decisions based on game state analysis
- Supports different difficulty levels and personalities
- Provides reasoning for AI actions

## Key Features
- **Turn-based gameplay** with phase management
- **Effect stacking system** with configurable rules
- **Event-driven architecture** for component communication
- **Comprehensive state tracking** with full audit trails
- **AI opponents** with different strategies and difficulties
- **Modular design** enabling easy testing and extension

## Development Commands

### Testing
```bash
# Run all tests
go test ./...

# Run tests for specific component
go test ./playerstate/...
go test ./game/...
```

### Build
```bash
# Build all components
go build ./...

# Build specific component
go build ./playerstate
```

### Lint
```bash
# Run linter on all code
golangci-lint run

# Run linter on specific component
golangci-lint run ./playerstate/...
```

## Architecture Principles
- **Domain-driven design** with clear component boundaries
- **Event-driven communication** between components
- **Interface-based design** for testability and extensibility
- **Immutable state changes** with atomic operations
- **Comprehensive logging** for debugging and replay
- **Clean separation** between game logic and infrastructure

## Implementation Status
✅ Overall architecture designed
✅ Package structure created
✅ Core types defined for all components
✅ Player state component fully stubbed with detailed comments
✅ Integration patterns established
✅ Documentation structure created
⏳ Implementation of stub functions (TODO)
⏳ Unit tests for all components (TODO)
⏳ Integration testing (TODO)
⏳ AI implementation (TODO)
⏳ Network/multiplayer support (TODO)

## Next Steps
1. **Implement player state stub functions** - Core functionality for state management
2. **Create comprehensive unit tests** - Ensure component reliability
3. **Implement game engine** - Turn management and rule enforcement
4. **Build card system** - Card definitions and effect processing
5. **Add combat mechanics** - Battle resolution and damage calculation
6. **Develop AI opponents** - Computer player decision making
7. **Add persistence layer** - Save/load game state
8. **Implement multiplayer** - Network protocol and state synchronization