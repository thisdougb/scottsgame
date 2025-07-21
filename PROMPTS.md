
## Prompts

These are the prompts I used to generate this repo content.

#### Player state

I asked just for the player state.

```
## goal

  I would like to plan the player state component of a dungeons and dragons style turn-based card game, written in Go.

  ## context

  - State tracks health, magic spells active, armor, etc. The common character and character modifiers in such a game.
  - The state component is distinct from other components in the architecture (such as inventory manager, game state, etc)
  - Game cards (spell cards, monster cards, condition cards, etc) can change state.
  - State changes stack up (for example health spell + armoured cloak may overlap)
  - State changes need to be tracked by turn, so we can expire them at the correct turn

  ## task

   - create package architecture for state
   - create state types required
   - create stub functions for principle code paths
   - create CLAUDE.md and ARCHITECTURE.md
   - create comments on the stub functions that describe their purpose
   - create a clear API into the state package
```

#### Game architecture

That looked a bit lonely, so I think asked Claude to add the whole game architecture.

```
can you move state specific doc into the state package?  I should have been clearer.  create a top level ARCHITECTURE.md that describes the game architecure, including our player state design.
```

