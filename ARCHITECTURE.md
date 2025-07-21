# D&D Turn-Based Card Game Architecture

## Overview
This is a dungeons and dragons style turn-based card game built in Go. The architecture follows domain-driven design principles with clear separation of concerns between game mechanics, player state, card systems, and supporting infrastructure.

## System Architecture

### Core Design Principles
- **Domain-Driven Design**: Each package represents a distinct domain concept
- **Event-Driven Architecture**: Components communicate through events
- **Clean Architecture**: Dependencies point inward toward business logic
- **Testability**: Interfaces enable easy mocking and unit testing
- **Immutability**: State changes are atomic and logged

### Component Overview
```
┌─────────────────────────────────────────────────────────┐
│                    Game Engine                          │
├─────────────┬─────────────┬─────────────┬───────────────┤
│   Player    │    Cards    │   Combat    │   Inventory   │
│   State     │   System    │   System    │   System      │
├─────────────┼─────────────┼─────────────┼───────────────┤
│   Events    │     AI      │   Network   │ Persistence   │
│   System    │   Players   │  Protocol   │    Layer      │
└─────────────┴─────────────┴─────────────┴───────────────┘
```

## Package Structure

### Core Game Packages

#### `game/` - Game State & Turn Management
- **Purpose**: Manages overall game state, turn order, and game phases
- **Key Types**: `Game`, `TurnAction`, `GamePhase`, `GameStatus`
- **Responsibilities**:
  - Turn order and phase management
  - Game state transitions
  - Win/loss condition checking
  - Action validation and coordination

#### `playerstate/` - Player Attributes & Effects
- **Purpose**: Tracks dynamic player attributes (health, mana, armor) and effects
- **Key Types**: `PlayerState`, `Effect`, `Modifier`, `StateChange`
- **Responsibilities**:
  - Health/mana/armor tracking
  - Temporary effects with turn-based expiration
  - Permanent modifiers from equipment/spells  
  - State change history and validation
- **Documentation**: See `playerstate/README.md` for detailed design

#### `cards/` - Card System
- **Purpose**: Defines cards, their effects, and deck management
- **Key Types**: `Card`, `Deck`, `Hand`, `CardEffect`
- **Responsibilities**:
  - Card definitions and metadata
  - Deck construction and shuffling
  - Hand management
  - Card effect definitions

#### `combat/` - Battle Mechanics
- **Purpose**: Handles combat resolution and damage calculation
- **Key Types**: `Combat`, `Attack`, `DamageCalculation`
- **Responsibilities**:
  - Combat phase management
  - Damage calculation with armor/resistance
  - Attack resolution and targeting
  - Combat event generation

#### `inventory/` - Items & Equipment
- **Purpose**: Manages player items and equipment
- **Key Types**: `Inventory`, `Item`, `ItemEffect`
- **Responsibilities**:
  - Item storage and management
  - Equipment slot management
  - Item effects and stat bonuses
  - Inventory capacity limits

### Supporting Systems

#### `events/` - Event System
- **Purpose**: Provides pub/sub messaging between components
- **Key Types**: `Event`, `EventSystem`, `GameLog`
- **Responsibilities**:
  - Event publishing and subscription
  - Game action logging
  - Cross-component communication
  - Replay functionality support

#### `ai/` - Computer Players
- **Purpose**: Implements AI opponents with different strategies
- **Key Types**: `AI`, `Decision`, `Strategy`
- **Responsibilities**:
  - AI decision making
  - Strategy implementation
  - Difficulty scaling
  - Performance evaluation

### Future Extension Points

#### `network/` - Multiplayer Support
- Real-time multiplayer protocol
- State synchronization
- Connection management
- Anti-cheat validation

#### `persistence/` - Save/Load System  
- Game state serialization
- Player progress tracking
- Deck storage
- Settings management

#### `ui/` - User Interface
- Game board rendering
- Card display and interaction
- Player status displays
- Menu systems

## Data Flow

### Turn-Based Game Loop
1. **Turn Start**: 
   - Events system notifies all components
   - Player state processes beginning-of-turn effects
   - AI makes decisions for computer players

2. **Player Actions**:
   - Cards played through card system
   - Combat initiated through combat system
   - Items used through inventory system
   - All actions validated by game engine

3. **Effect Resolution**:
   - Player state calculates effective values
   - Combat system resolves damage
   - Events logged for all changes

4. **Turn End**:
   - Player state processes end-of-turn effects
   - Effects expire, modifiers update
   - Game checks win/loss conditions

### Inter-Component Communication

#### Event-Driven Updates
```
Card Played → Events → Player State → Combat → Events → UI Update
     ↓                     ↓              ↓
  Game Log         State Changes    Damage Calc
```

#### State Queries
- UI queries player state for display
- AI queries all systems for decision making
- Combat system queries player state for damage calculation
- Game engine queries all systems for validation

## Integration Patterns

### Dependency Injection
- All major components define interfaces
- Concrete implementations injected at runtime
- Enables easy testing with mocks
- Supports different implementations (local vs network)

### Event Sourcing
- All state changes generate events
- Events stored for replay/debugging
- Enables rollback for network issues
- Supports spectator mode

### Command Pattern
- Player actions encapsulated as commands
- Commands validated before execution
- Supports undo functionality
- Network serialization friendly

## Error Handling Strategy

### Validation Layers
1. **Input Validation**: UI/API boundary validation
2. **Business Rules**: Game rule enforcement  
3. **State Consistency**: Cross-component validation
4. **System Integrity**: Internal consistency checks

### Recovery Mechanisms
- Failed actions don't change game state
- Event log enables state reconstruction
- Graceful degradation for non-critical failures
- Network reconnection and state synchronization

## Performance Considerations

### Memory Management
- Object pooling for frequently created objects
- Efficient collections for large card sets
- Minimal allocation in hot paths
- Memory-mapped files for large assets

### Concurrency
- Read-heavy operations use concurrent-safe collections
- Write operations use channels for coordination
- AI computation runs in separate goroutines
- Network I/O uses async patterns

### Scalability
- Stateless components enable horizontal scaling
- Event sourcing supports distributed systems
- Component isolation enables selective optimization
- Caching layers for expensive calculations

## Testing Strategy

### Unit Testing
- Each package has comprehensive unit tests
- Interface-based mocking for dependencies
- Property-based testing for game rules
- Benchmark tests for performance-critical paths

### Integration Testing  
- Full game scenarios with multiple players
- AI vs AI automated testing
- Network protocol testing
- State consistency validation

### System Testing
- Load testing with many concurrent games
- Chaos engineering for failure scenarios
- Performance profiling and optimization
- Security testing for multiplayer features