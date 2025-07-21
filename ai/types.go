package ai

// AI represents an artificial intelligence player
type AI struct {
	ID         string
	PlayerID   string
	Difficulty Difficulty
	Strategy   Strategy
	Personality AIPersonality
}

// Difficulty defines AI skill levels
type Difficulty int

const (
	DifficultyEasy Difficulty = iota
	DifficultyMedium
	DifficultyHard
	DifficultyExpert
)

// Strategy defines AI playing styles
type Strategy int

const (
	StrategyAggressive Strategy = iota
	StrategyDefensive
	StrategyBalanced
	StrategyControl
	StrategyRush
)

// AIPersonality affects decision making
type AIPersonality int

const (
	PersonalityCalculated AIPersonality = iota
	PersonalityRisky
	PersonalityConservative
	PersonalityUnpredictable
)

// Decision represents an AI decision with reasoning
type Decision struct {
	Action      ActionType
	Confidence  float64
	Reasoning   string
	Alternatives []Alternative
}

// Alternative represents other options the AI considered
type Alternative struct {
	Action ActionType
	Score  float64
	Reason string
}