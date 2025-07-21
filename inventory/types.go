package inventory

// Inventory manages items and equipment for players
type Inventory struct {
	PlayerID  string
	Items     []Item
	Equipment map[SlotType]Item
	Capacity  int
}

// Item represents any item that can be owned by a player
type Item struct {
	ID          string
	Name        string
	Description string
	Type        ItemType
	Rarity      ItemRarity
	Value       int
	Stackable   bool
	MaxStack    int
	
	// Equipment properties
	SlotType   SlotType
	StatBonuses map[string]int
	Effects     []ItemEffect
}

// ItemType categorizes different item types
type ItemType int

const (
	ItemTypeWeapon ItemType = iota
	ItemTypeArmor
	ItemTypeAccessory
	ItemTypeConsumable
	ItemTypeMaterial
	ItemTypeQuest
)

// SlotType defines equipment slots
type SlotType int

const (
	SlotTypeWeapon SlotType = iota
	SlotTypeArmor
	SlotTypeHelmet
	SlotTypeBoots
	SlotTypeRing
	SlotTypeNecklace
)

// ItemRarity defines item rarity levels
type ItemRarity int

const (
	ItemRarityCommon ItemRarity = iota
	ItemRarityUncommon
	ItemRarityRare
	ItemRarityEpic
	ItemRarityLegendary
)

// ItemEffect represents effects that items can provide
type ItemEffect struct {
	Type     EffectType
	Value    int
	Duration int // -1 for permanent while equipped
}