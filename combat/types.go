package combat

// Combat manages battle mechanics and damage resolution
type Combat struct {
	GameID      string
	Participants []string
	Round       int
	Phase       CombatPhase
	Resolved    bool
}

// CombatPhase represents phases of combat resolution
type CombatPhase int

const (
	PhaseDeclaration CombatPhase = iota
	PhaseTargeting
	PhaseResolution
	PhaseCleanup
)

// Attack represents a combat action between entities
type Attack struct {
	AttackerID string
	DefenderID string
	Damage     int
	AttackType AttackType
	Resolved   bool
}

// AttackType defines different types of attacks
type AttackType int

const (
	AttackTypeMelee AttackType = iota
	AttackTypeRanged
	AttackTypeMagical
	AttackTypeSpell
)

// DamageCalculation holds the breakdown of damage calculation
type DamageCalculation struct {
	BaseDamage      int
	Modifiers       int
	ArmorReduction  int
	FinalDamage     int
	DamageType      DamageType
}

// DamageType categorizes different damage types
type DamageType int

const (
	DamageTypePhysical DamageType = iota
	DamageTypeFire
	DamageTypeIce
	DamageTypeLightning
	DamageTypePoison
	DamageTypeHealing
)