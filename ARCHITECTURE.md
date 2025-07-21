# Player State Component Architecture

## Overview
The player state component manages all dynamic attributes of player characters in the D&D-style card game. It handles health, mana, armor, temporary effects, and permanent modifiers while maintaining turn-based consistency.

## Design Principles

### Separation of Concerns
- **State Storage**: Core data structures for player attributes
- **State Management**: Business logic for state transitions  
- **Effect System**: Turn-based temporary modifications
- **Modifier System**: Persistent attribute changes
- **API Layer**: Clean interface for external systems

### Immutability and Consistency
- State changes are atomic operations
- All changes are logged for audit trails
- Validation ensures state remains within valid bounds
- Turn processing is deterministic

## Component Breakdown

### Core Types (`types.go`)
```go
PlayerState     // Main state container
Effect          // Temporary turn-based changes  
Modifier        // Persistent attribute changes
StateChange     // Change history record
```

**Key Design Decisions:**
- Effects have duration in turns and expiration logic
- Modifiers can be permanent or temporary 
- Both effects and modifiers track their source
- State changes are immutable records

### State Manager (`manager.go`) 
The central coordinator handling:
- Player lifecycle (create, retrieve, validate)
- Core operations (damage, healing, resource consumption)
- Turn processing (start/end turn effects)
- State calculation with modifiers applied

**Key Patterns:**
- Manager maintains player registry and change history
- All operations return errors for proper error handling
- State changes are logged immediately when applied
- Turn processing is event-driven

### Effects System (`effects.go`)
Handles temporary modifications that expire:
- Adding/removing effects with duration tracking
- Stacking rules for multiple effects of same type
- Turn-based duration management
- Event triggers (on damage, on heal, etc.)

**Stacking Strategy:**
- Stackable effects accumulate (armor bonuses)
- Non-stackable effects replace (shields take highest)
- Source tracking prevents duplicate effects from same card

### Modifiers System (`modifiers.go`)
Manages longer-term stat changes:
- Equipment bonuses (armor, weapons)
- Permanent spell effects
- Character advancement bonuses
- Calculated attribute values

**Modifier Types:**
- Additive: Simple number additions/subtractions
- Multiplicative: Percentage-based modifications  
- Override: Replace base value entirely

### Public API (`api.go`)
Clean interface for external systems:
- StateManager interface for dependency injection
- CardEffect integration for game card system
- Batch operations for complex multi-step changes
- UI-friendly summary data structures

## Data Flow

### Turn Processing
1. **Turn Start**: Process beginning-of-turn effects (regeneration, poison)
2. **Card Play**: Apply card effects through API
3. **State Calculation**: Compute effective values with all modifiers  
4. **Turn End**: Process end-of-turn effects, expire durations
5. **Cleanup**: Remove expired effects, validate state

### Card Application
1. Parse CardEffect definition
2. Validate targeting rules  
3. Apply direct stat changes
4. Add effects and modifiers
5. Log all changes with card as source
6. Return success/failure to game engine

### State Queries
1. Retrieve base player state
2. Apply all active modifiers
3. Apply all active effects  
4. Return calculated effective values
5. Cache results for performance

## Integration Points

### Game Engine Integration
- StateManager interface allows easy mocking for tests
- Card system calls ApplyCardEffect for all card plays
- Turn system calls ProcessTurnStart/End for each player
- UI queries GetPlayerSummary for display

### Future Extensions
- **Persistence Layer**: Save/load player states
- **Network Layer**: Synchronize state across clients  
- **Analytics**: Track state changes for balancing
- **AI Integration**: State evaluation for computer players

## Error Handling Strategy

### Validation Levels
1. **Input Validation**: Check parameters at API boundaries
2. **State Validation**: Ensure state consistency after changes
3. **Business Logic**: Validate game rules (sufficient mana, etc.)
4. **System Validation**: Verify internal data consistency

### Recovery Patterns
- Failed operations don't modify state
- State history enables rollback if needed
- Validation errors include details for debugging
- Graceful degradation for non-critical failures

## Performance Considerations

### Optimization Strategies
- Cache effective attribute calculations
- Lazy evaluation of complex state queries
- Batch operations for multiple simultaneous changes
- Minimal memory allocation in hot paths

### Scalability Notes  
- State is kept in memory for fast access
- History pruning for long-running games
- Concurrent access patterns need synchronization
- Consider event sourcing for distributed systems

## Testing Strategy

### Unit Test Coverage
- Each function has positive and negative test cases
- Effect stacking logic thoroughly tested
- Turn processing edge cases covered
- State validation boundary conditions tested

### Integration Testing
- Full turn cycles with multiple players
- Complex card effect combinations
- Error recovery scenarios
- Performance testing with large effect counts